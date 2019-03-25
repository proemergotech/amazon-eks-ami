Read the original README for ami related information.  
https://github.com/awslabs/amazon-eks-ami

To update the `eni-max-pods.txt`, run `update_ip_limit.go`.  
We need to provide a `subnet_id` to make because we don't have a default vpc.
```bash
make SUBNET_ID=subnet-11111
```  
It can be any subnet with internet access, but prefer one from the `dev` environment.  

