package main

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/ec2"
	"github.com/olekukonko/tablewriter"
	"os"
)

func main() {

	// create table to better parse output
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Name", "Instance ID", "Instance Type", "AMI", "Status"})

	svc := ec2.New(&aws.Config{Region: aws.String("eu-west-1")})

	resp, err := svc.DescribeInstances(nil)
	if err != nil {
		panic(err)
	}

	// go over list of instances and print them nicely on an ASCII table
	for _, reservations := range resp.Reservations {
		for _, inst := range reservations.Instances {
			// go through list of tags and get instanceName into name
			var instanceName string
			for _, tag := range inst.Tags {
				if *tag.Key == "Name" {
					instanceName = *tag.Value
				}
			}

			// add values under table on Stdout provided by OS package
			table.Append([]string{
				instanceName,
				*inst.InstanceId,
				*inst.InstanceType,
				*inst.ImageId,
				*inst.State.Name, // yes! it requires a ',' as weirdo as this might be
			})
		}
	}
	// happliy pretty print table
	table.Render()
}
