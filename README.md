# Sleep all your Okteto namespaces on a schedule

> This is an experiment and Okteto does not officially support it.

- Create an [Okteto Admin Token](https://www.okteto.com/docs/admin/dashboard/#admin-access-tokens)

- Export the token to a local variable:

```bash
export OKTETO_ADMIN_TOKEN=<<your-token>>
```

- Create a namespace, and, via the admin section, mark it as [Keep awake](https://www.okteto.com/docs/admin/dashboard/#namespaces)

- Export the namespace name to a local variable:

```bash
export NAMESPACE=<<your-namespace>>
```

- Create a local variable to define the sleep cronjob schedule:

```bash
export SLEEP_JOB_SCHEDULE="0 20 * * *"
```

For example, 0 0 13 * 5 states that the task must be started every Friday at midnight, as well as on the 13th of each month at midnight.

- Run the following command to create the cronjob:

```bash
okteto deploy -n ${NAMESPACE} --var OKTETO_ADMIN_TOKEN=${OKTETO_ADMIN_TOKEN} --var SLEEP_JOB_SCHEDULE=${SLEEP_JOB_SCHEDULE}
```

## Force the execution of the job

To force the execution of the sleep namespaces job, run the following commands:

```bash
okteto kubeconfig
kubectl -n ${NAMESPACE} create job --from=cronjob/sleep-all-namespaces sleep-all-namespaces-$(date +%s)
```

## Golang code

This sample is written in golang, and it uses [openapi-generator](https://openapi-generator.tech/) to generate the client of the API from the Okteto API specification in your instance. You can locate the
Open API specification in your Okteto instance at `https://<your-okteto-instance>/api/schema/v0/openapi.json `.