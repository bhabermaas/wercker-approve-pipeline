# wercker-approve-pipeline

This is a simple demonstration of the Wercker API to trigger a pipeline that
has been marked as manual requiring approval. 

Given the following command arguments - 
username application pipeline-name commitHash wercker-token

It will determine the applicationID and get all runs for a specific commitHash. 
The list of runs will be inspected for the specified pipeline that requires approval. 
The pipeline will be triggered and then an approval request performed. 