Read the original README for ami related information.  
https://github.com/awslabs/amazon-eks-ami

To update the `eni-max-pods.txt`, run `update_ip_limit.go`.  
Then change the `subnet_id` in `eks-worker-al2.json` variables section if necessary.  
We need to provide this variable because we don't have a default vpc.  