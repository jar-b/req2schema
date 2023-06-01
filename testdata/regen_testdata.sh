#!/bin/bash

prefixes="batch_job_queue ecs_cluster ecs_service guardduty_findings medialive_channel medialive_multiplex simple"

for prefix in $prefixes; do
    req2schema -out=testdata/${prefix}.go testdata/${prefix}.json
done

