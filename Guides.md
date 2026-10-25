# Guides

## Enable dev environment to generate Signed URLs

gcloud auth application-default login --impersonate-service-account=andres-castro-photography-app@marine-duality-444902-n2.iam.gserviceaccount.com
This creates local Application Default Credentials (ADC) configured to impersonate the service account. Google's Go client libraries support this approach.
