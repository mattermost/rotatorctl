package main

import (
	rotatorModel "github.com/mattermost/rotator/model"
	"github.com/mattermost/rotator/rotator"
	"github.com/pkg/errors"
	logger "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

type rotateOptions struct {
	cluster              string
	maxScaling           int
	rotateMasters        bool
	rotateWorkers        bool
	maxDrainRetries      int
	evictGracePeriod     int
	waitBetweenRotations int
	waitBetweenDrains    int
}

func newRotateCmd() *cobra.Command {
	o := rotateOptions{}

	cmd := &cobra.Command{
		Use:   "rotate",
		Short: "Rotates nodes",
		Long:  `Applying latest AMI that exists on launch template or launch configuration and rolls out new nodes`,
		RunE: func(cmd *cobra.Command, args []string) error {
			logger := logger.WithField("cluster", o.cluster)

			clusterRotator := rotatorModel.Cluster{
				ClusterID:            o.cluster,
				MaxScaling:           o.maxScaling,
				RotateMasters:        o.rotateMasters,
				RotateWorkers:        o.rotateWorkers,
				MaxDrainRetries:      o.maxDrainRetries,
				EvictGracePeriod:     o.evictGracePeriod,
				WaitBetweenRotations: o.waitBetweenRotations,
				WaitBetweenDrains:    o.waitBetweenDrains,
			}
			rotatorMetadata, err := rotator.InitRotateCluster(&clusterRotator, &rotator.RotatorMetadata{}, logger)
			if err != nil {
				return err
			}
			if err = printJSON(rotatorMetadata); err != nil {
				return errors.Wrap(err, "failed to print cluster response")
			}

			return nil
		},
	}

	cmd.Flags().StringVar(&o.cluster, "cluster", "", "The name of the cluster to be upgraded.")
	cmd.Flags().IntVar(&o.maxScaling, "max-scaling", 5, "The maximum number of nodes to rotate every time. If the number is bigger than the number of nodes, then the number of nodes will be the maximum number.")
	cmd.Flags().BoolVar(&o.rotateMasters, "rotate-masters", false, "if disabled, master nodes will not be rotated")
	cmd.Flags().BoolVar(&o.rotateWorkers, "rotate-workers", true, "if disabled, worker nodes will not be rotated")
	cmd.Flags().IntVar(&o.maxDrainRetries, "max-drain-retries", 10, "The number of times to retry a node drain.")
	cmd.Flags().IntVar(&o.evictGracePeriod, "evict-grace-period", 600, "The pod eviction grace period when draining in seconds.")
	cmd.Flags().IntVar(&o.waitBetweenRotations, "wait-between-rotations", 60, "Τhe time to wait between each rotation of a group of nodes.")
	cmd.Flags().IntVar(&o.waitBetweenDrains, "wait-between-drains", 60, "The time to wait between each node drain in a group of nodes.")
	cmd.MarkFlagRequired("cluster")

	return cmd
}
