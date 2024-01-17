# Kubernetes Cluster Setup

This README guides you through the process of setting up a Kubernetes cluster with vms using Vagrant and Ansible. Ensure that you have all necessary dependencies installed, including Ansible, Vagrant, and any SSH tools needed.

## Steps to Follow

1. First, change to the `kubernetes-cluster` directory
2. Move your ssh public key into `kubernetes-cluster/ubuntu/vagrant/id_rsa.pub`
3. Start the vm deployment with `vagrant up`
4. Change, to the `ansible_playbooks` folder
5. Optional: Sometimes you have to specify the ansible config
    ```shell
    export ANSIBLE_CONFIG=./ansible.cfg
    ```
6. Run `ansible-playbook setup.yaml` to setup cluster

> The first time the manifests are automatically deployed after commissioning, the distribution of the pods may not be uniform.

## Additional Notes
Ensure that Ansible is properly installed and configured on your system.
Verify that the SSH key has the correct permissions and is recognized by your system.
If you encounter any issues, consult the official Kubernetes and Ansible documentation for troubleshooting steps.
Conclusion
Following these steps should successfully set up a Kubernetes cluster using Ansible. For more detailed instructions or advanced configurations, refer to the Kubernetes and Ansible official documentation.