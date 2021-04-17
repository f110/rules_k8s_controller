package main

import (
	"bufio"
	"bytes"
	"context"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"time"

	"github.com/spf13/cobra"
	"golang.org/x/xerrors"
	goyaml "gopkg.in/yaml.v2"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	configv1alpha4 "sigs.k8s.io/kind/pkg/apis/config/v1alpha4"
)

func main() {
	rootCmd := &cobra.Command{
		Use: "kindcluster",
	}
	createCmd(rootCmd)
	deleteCmd(rootCmd)

	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "%+v\n", err)
		os.Exit(1)
	}
}

func createCmd(rootCmd *cobra.Command) {
	var name, kind, k8sVersion string
	workerNum := 1
	cmd := &cobra.Command{
		Use: "create",
		RunE: func(_ *cobra.Command, _ []string) error {
			cluster, err := NewCluster(kind, name, "")
			if err != nil {
				return xerrors.Errorf(": %w", err)
			}
			if err := cluster.Create(k8sVersion, workerNum); err != nil {
				return xerrors.Errorf(": %w", err)
			}

			return nil
		},
	}
	cmd.Flags().StringVar(&name, "name", "kindcluster", "Name of cluster")
	cmd.Flags().StringVar(&kind, "kind", "", "The path of kind")
	cmd.Flags().StringVar(&k8sVersion, "k8s-version", "v1.20.2", "Cluster version")
	cmd.Flags().IntVar(&workerNum, "worker-num", workerNum, "The number of worker")
	// TODO:
	//cmd.Flags().StringVar(&manifest, "manifest", "", "The path of default manifest")

	rootCmd.AddCommand(cmd)
}

func deleteCmd(rootCmd *cobra.Command) {
	var name, kind string
	cmd := &cobra.Command{
		Use: "delete",
		RunE: func(_ *cobra.Command, _ []string) error {
			cluster, err := NewCluster(kind, name, "")
			if err != nil {
				return xerrors.Errorf(": %w", err)
			}
			if err := cluster.Delete(); err != nil {
				return xerrors.Errorf(": %w", err)
			}
			return nil
		},
	}
	cmd.Flags().StringVar(&name, "name", "kindcluster", "Name of cluster")
	cmd.Flags().StringVar(&kind, "kind", "", "The path of kind")

	rootCmd.AddCommand(cmd)
}

var KindNodeImageHash = map[string]string{
	"v1.20.2":  "8f7ea6e7642c0da54f04a7ee10431549c0257315b3a634f6ef2fecaaedb19bab",
	"v1.19.7":  "a70639454e97a4b733f9d9b67e12c01f6b0297449d5b9cbbef87473458e26dca",
	"v1.19.3":  "e1ac015e061da4b931cc4f693e22d7bc1110f031faf7b2af4c4fefac9e65565d",
	"v1.19.1":  "98cf5288864662e37115e362b23e4369c8c4a408f99cbc06e58ac30ddc721600",
	"v1.19.0":  "3b0289b2d1bab2cb9108645a006939d2f447a10ad2bb21919c332d06b548bbc6",
	"v1.18.15": "5c1b980c4d0e0e8e7eb9f36f7df525d079a96169c8a8f20d8bd108c0d0889cc4",
	"v1.18.8":  "f4bcc97a0ad6e7abaf3f643d890add7efe6ee4ab90baeb374b4f41a4c95567eb",
}

type Cluster struct {
	kind          string
	name          string
	kubeConfig    string
	tmpKubeConfig bool

	clientset kubernetes.Interface
}

func NewCluster(kind, name, kubeConfig string) (*Cluster, error) {
	_, err := exec.LookPath(kind)
	if err != nil {
		return nil, err
	}

	return &Cluster{kind: kind, name: name, kubeConfig: kubeConfig}, nil
}

func (c *Cluster) IsExist(name string) (bool, error) {
	cmd := exec.CommandContext(context.TODO(), c.kind, "get", "clusters")
	buf, err := cmd.CombinedOutput()
	if err != nil {
		return false, xerrors.Errorf(": %w", err)
	}
	s := bufio.NewScanner(bytes.NewReader(buf))
	for s.Scan() {
		line := s.Text()
		if line == name {
			return true, nil
		}
	}

	return false, nil
}

func (c *Cluster) Create(clusterVersion string, workerNum int) error {
	kindConfFile, err := os.CreateTemp("", "kind.config.yaml")
	if err != nil {
		return xerrors.Errorf(": %w", err)
	}
	defer os.Remove(kindConfFile.Name())

	imageHash, ok := KindNodeImageHash[clusterVersion]
	if !ok {
		return xerrors.Errorf("Not supported k8s version: %s", clusterVersion)
	}
	image := fmt.Sprintf("kindest/node:%s@sha256:%s", clusterVersion, imageHash)

	clusterConf := &configv1alpha4.Cluster{
		TypeMeta: configv1alpha4.TypeMeta{
			APIVersion: "kind.x-k8s.io/v1alpha4",
			Kind:       "Cluster",
		},
		Nodes: []configv1alpha4.Node{
			{Role: configv1alpha4.ControlPlaneRole, Image: image},
		},
	}
	// If workerNum equals 1 is intended to create a single node cluster.
	// In that case, We shouldn't add Node.
	if workerNum > 2 {
		for i := 0; i < workerNum; i++ {
			clusterConf.Nodes = append(clusterConf.Nodes,
				configv1alpha4.Node{Role: configv1alpha4.WorkerRole, Image: image})
		}
	}
	if buf, err := goyaml.Marshal(clusterConf); err != nil {
		return xerrors.Errorf(": %w", err)
	} else {
		if _, err := kindConfFile.Write(buf); err != nil {
			return xerrors.Errorf(": %w", err)
		}
	}

	if c.kubeConfig == "" {
		f, err := os.CreateTemp("", "config")
		if err != nil {
			return err
		}
		c.kubeConfig = f.Name()
		c.tmpKubeConfig = true
	}
	cmd := exec.CommandContext(
		context.TODO(),
		c.kind, "create", "cluster",
		"--name", c.name,
		"--kubeconfig", c.kubeConfig,
		"--config", kindConfFile.Name(),
	)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		return err
	}

	return nil
}

func (c *Cluster) KubeConfig() string {
	return c.kubeConfig
}

func (c *Cluster) Delete() error {
	found, err := c.IsExist(c.name)
	if err != nil {
		return xerrors.Errorf(": %w", err)
	}
	if !found {
		return nil
	}

	if c.tmpKubeConfig {
		defer os.Remove(c.kubeConfig)
	}
	cmd := exec.CommandContext(context.TODO(), c.kind, "delete", "cluster", "--name", c.name)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

func (c *Cluster) RESTConfig() (*rest.Config, error) {
	if exist, err := c.IsExist(c.name); err != nil {
		return nil, xerrors.Errorf(": %w", err)
	} else if !exist {
		return nil, xerrors.New("The cluster is not created yet")
	}
	if c.kubeConfig == "" {
		kubeConf, err := os.CreateTemp("", "kubeconfig")
		if err != nil {
			return nil, xerrors.Errorf(": %w", err)
		}
		cmd := exec.CommandContext(
			context.TODO(),
			c.kind, "export", "kubeconfig",
			"--kubeconfig", kubeConf.Name(),
			"--name", c.name,
		)
		if err := cmd.Run(); err != nil {
			return nil, xerrors.Errorf(": %w", err)
		}
		c.kubeConfig = kubeConf.Name()
		defer func() {
			os.Remove(kubeConf.Name())
			c.kubeConfig = ""
		}()
	}

	cfg, err := clientcmd.LoadFromFile(c.kubeConfig)
	if err != nil {
		return nil, xerrors.Errorf(": %w", err)
	}
	clientConfig := clientcmd.NewDefaultClientConfig(*cfg, &clientcmd.ConfigOverrides{})
	restCfg, err := clientConfig.ClientConfig()
	if err != nil {
		return nil, err
	}

	return restCfg, nil
}

func (c *Cluster) Clientset() (kubernetes.Interface, error) {
	if c.clientset != nil {
		return c.clientset, nil
	}

	restCfg, err := c.RESTConfig()
	if err != nil {
		return nil, xerrors.Errorf(": %w", err)
	}
	cs, err := kubernetes.NewForConfig(restCfg)
	if err != nil {
		return nil, err
	}
	c.clientset = cs

	return cs, nil
}

func (c *Cluster) WaitReady(ctx context.Context) error {
	client, err := c.Clientset()
	if err != nil {
		return xerrors.Errorf(": %w", err)
	}

	return PollImmediate(ctx, 1*time.Second, 3*time.Minute, func(ctx context.Context) (done bool, err error) {
		nodes, err := client.CoreV1().Nodes().List(ctx, metav1.ListOptions{})
		if err != nil {
			return false, err
		}

		notReadyNodes := make(map[string]struct{})
	Nodes:
		for _, v := range nodes.Items {
			for _, c := range v.Status.Conditions {
				if c.Type == corev1.NodeReady && c.Status == corev1.ConditionTrue {
					continue Nodes
				}
			}
			notReadyNodes[v.Name] = struct{}{}
		}
		if len(notReadyNodes) == 0 {
			return true, nil
		}

		return false, nil
	})
}

func Poll(ctx context.Context, interval, timeout time.Duration, fn func(ctx context.Context) (done bool, err error)) error {
	tick := time.NewTicker(interval)
	defer tick.Stop()

	limit := time.After(timeout)
	for {
		select {
		case <-tick.C:
			fnCtx, cancel := context.WithTimeout(ctx, interval)
			done, err := fn(fnCtx)
			cancel()
			if err != nil {
				return xerrors.Errorf(": %w", err)
			}
			if done {
				return nil
			}
		case <-limit:
			return errors.New("timed out")
		case <-ctx.Done():
			return ctx.Err()
		}
	}
}

func PollImmediate(ctx context.Context, interval, timeout time.Duration, fn func(ctx context.Context) (done bool, err error)) error {
	fnCtx, cancel := context.WithTimeout(ctx, interval)
	done, err := fn(fnCtx)
	cancel()
	if done {
		return nil
	}
	if err != nil {
		return xerrors.Errorf(": %w", err)
	}

	return Poll(ctx, interval, timeout, fn)
}
