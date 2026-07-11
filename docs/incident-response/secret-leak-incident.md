# Secret leak incident

## Severity classification
SEV-1 for active production/admin credential exposure; SEV-2 for scoped secret with plausible access; SEV-3 for expired/synthetic value.

## Symptoms
Scanner alert, credential in Git/log/state/artifact/chat, unexpected authentication or cloud audit activity.

## Initial checks
Privately identify secret type, scope, exposure window, repositories/artifacts and observed use. Do not copy the value.

## Containment
Revoke/disable first where safe, block affected pipeline/artifact access and issue a replacement through the owning secret system.

## Diagnosis
Trace creation, propagation, access logs and every derivative location including state/history/caches.

## Remediation
Rotate dependent credentials, update references, remove public artifacts/history using provider procedures and add detection/prevention.

## Rollback
Do not restore the old secret. Roll back application configuration only to a version referencing the new secret.

## Communication notes
Use a private security channel; share identifiers/fingerprints, not values. Coordinate disclosure and customer/legal review.

## Evidence collection
Preserve redacted scanner output, audit events, exposure timestamps, rotation IDs and affected commits/artifacts.

## Postmortem checklist
Rotation completeness; unauthorized use; root cause; history/cache cleanup; detector tests; owners/dates.
