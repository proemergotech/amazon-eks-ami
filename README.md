Read the original README for ami related information.  
https://github.com/awslabs/amazon-eks-ami

To update the `eni-max-pods.txt`, run `update_ip_limit.go`.  
Change the `ami_name` and `subnet_id` (if necessary) in `eks-worker-al2.json` variables section.  
We need to provide the `subnet_id` variable because we don't have a default vpc.  