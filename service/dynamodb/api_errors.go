// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package dynamodb

const (

	// ErrCodeBackupInUseException for service response error code
	// "BackupInUseException".
	//
	// There is another ongoing conflicting backup control plane operation on the
	// table. The backup is either being created, deleted or restored to a table.
	ErrCodeBackupInUseException = "BackupInUseException"

	// ErrCodeBackupNotFoundException for service response error code
	// "BackupNotFoundException".
	//
	// Backup not found for the given BackupARN.
	ErrCodeBackupNotFoundException = "BackupNotFoundException"

	// ErrCodeConditionalCheckFailedException for service response error code
	// "ConditionalCheckFailedException".
	//
	// A condition specified in the operation could not be evaluated.
	ErrCodeConditionalCheckFailedException = "ConditionalCheckFailedException"

	// ErrCodeContinuousBackupsUnavailableException for service response error code
	// "ContinuousBackupsUnavailableException".
	//
	// Backups have not yet been enabled for this table.
	ErrCodeContinuousBackupsUnavailableException = "ContinuousBackupsUnavailableException"

	// ErrCodeGlobalTableAlreadyExistsException for service response error code
	// "GlobalTableAlreadyExistsException".
	//
	// The specified global table already exists.
	ErrCodeGlobalTableAlreadyExistsException = "GlobalTableAlreadyExistsException"

	// ErrCodeGlobalTableNotFoundException for service response error code
	// "GlobalTableNotFoundException".
	//
	// The specified global table does not exist.
	ErrCodeGlobalTableNotFoundException = "GlobalTableNotFoundException"

	// ErrCodeIdempotentParameterMismatchException for service response error code
	// "IdempotentParameterMismatchException".
	//
	// DynamoDB rejected the request because you retried a request with a different
	// payload but with an idempotent token that was already used.
	ErrCodeIdempotentParameterMismatchException = "IdempotentParameterMismatchException"

	// ErrCodeIndexNotFoundException for service response error code
	// "IndexNotFoundException".
	//
	// The operation tried to access a nonexistent index.
	ErrCodeIndexNotFoundException = "IndexNotFoundException"

	// ErrCodeInternalServerError for service response error code
	// "InternalServerError".
	//
	// An error occurred on the server side.
	ErrCodeInternalServerError = "InternalServerError"

	// ErrCodeInvalidRestoreTimeException for service response error code
	// "InvalidRestoreTimeException".
	//
	// An invalid restore time was specified. RestoreDateTime must be between EarliestRestorableDateTime
	// and LatestRestorableDateTime.
	ErrCodeInvalidRestoreTimeException = "InvalidRestoreTimeException"

	// ErrCodeItemCollectionSizeLimitExceededException for service response error code
	// "ItemCollectionSizeLimitExceededException".
	//
	// An item collection is too large. This exception is only returned for tables
	// that have one or more local secondary indexes.
	ErrCodeItemCollectionSizeLimitExceededException = "ItemCollectionSizeLimitExceededException"

	// ErrCodeLimitExceededException for service response error code
	// "LimitExceededException".
	//
	// There is no limit to the number of daily on-demand backups that can be taken.
	//
	// Up to 50 simultaneous table operations are allowed per account. These operations
	// include CreateTable, UpdateTable, DeleteTable,UpdateTimeToLive, RestoreTableFromBackup,
	// and RestoreTableToPointInTime.
	//
	// The only exception is when you are creating a table with one or more secondary
	// indexes. You can have up to 25 such requests running at a time; however,
	// if the table or index specifications are complex, DynamoDB might temporarily
	// reduce the number of concurrent operations.
	//
	// There is a soft account limit of 256 tables.
	ErrCodeLimitExceededException = "LimitExceededException"

	// ErrCodePointInTimeRecoveryUnavailableException for service response error code
	// "PointInTimeRecoveryUnavailableException".
	//
	// Point in time recovery has not yet been enabled for this source table.
	ErrCodePointInTimeRecoveryUnavailableException = "PointInTimeRecoveryUnavailableException"

	// ErrCodeProvisionedThroughputExceededException for service response error code
	// "ProvisionedThroughputExceededException".
	//
	// Your request rate is too high. The AWS SDKs for DynamoDB automatically retry
	// requests that receive this exception. Your request is eventually successful,
	// unless your retry queue is too large to finish. Reduce the frequency of requests
	// and use exponential backoff. For more information, go to Error Retries and
	// Exponential Backoff (https://docs.aws.amazon.com/amazondynamodb/latest/developerguide/Programming.Errors.html#Programming.Errors.RetryAndBackoff)
	// in the Amazon DynamoDB Developer Guide.
	ErrCodeProvisionedThroughputExceededException = "ProvisionedThroughputExceededException"

	// ErrCodeReplicaAlreadyExistsException for service response error code
	// "ReplicaAlreadyExistsException".
	//
	// The specified replica is already part of the global table.
	ErrCodeReplicaAlreadyExistsException = "ReplicaAlreadyExistsException"

	// ErrCodeReplicaNotFoundException for service response error code
	// "ReplicaNotFoundException".
	//
	// The specified replica is no longer part of the global table.
	ErrCodeReplicaNotFoundException = "ReplicaNotFoundException"

	// ErrCodeRequestLimitExceeded for service response error code
	// "RequestLimitExceeded".
	//
	// Throughput exceeds the current throughput limit for your account. Please
	// contact AWS Support at AWS Support (https://docs.aws.amazon.com/https:/aws.amazon.com/support)
	// to request a limit increase.
	ErrCodeRequestLimitExceeded = "RequestLimitExceeded"

	// ErrCodeResourceInUseException for service response error code
	// "ResourceInUseException".
	//
	// The operation conflicts with the resource's availability. For example, you
	// attempted to recreate an existing table, or tried to delete a table currently
	// in the CREATING state.
	ErrCodeResourceInUseException = "ResourceInUseException"

	// ErrCodeResourceNotFoundException for service response error code
	// "ResourceNotFoundException".
	//
	// The operation tried to access a nonexistent table or index. The resource
	// might not be specified correctly, or its status might not be ACTIVE.
	ErrCodeResourceNotFoundException = "ResourceNotFoundException"

	// ErrCodeTableAlreadyExistsException for service response error code
	// "TableAlreadyExistsException".
	//
	// A target table with the specified name already exists.
	ErrCodeTableAlreadyExistsException = "TableAlreadyExistsException"

	// ErrCodeTableInUseException for service response error code
	// "TableInUseException".
	//
	// A target table with the specified name is either being created or deleted.
	ErrCodeTableInUseException = "TableInUseException"

	// ErrCodeTableNotFoundException for service response error code
	// "TableNotFoundException".
	//
	// A source table with the name TableName does not currently exist within the
	// subscriber's account.
	ErrCodeTableNotFoundException = "TableNotFoundException"

	// ErrCodeTransactionCanceledException for service response error code
	// "TransactionCanceledException".
	//
	// The entire transaction request was rejected.
	//
	// DynamoDB rejects a TransactWriteItems request under the following circumstances:
	//
	//    * A condition in one of the condition expressions is not met.
	//
	//    * A table in the TransactWriteItems request is in a different account
	//    or region.
	//
	//    * More than one action in the TransactWriteItems operation targets the
	//    same item.
	//
	//    * There is insufficient provisioned capacity for the transaction to be
	//    completed.
	//
	//    * An item size becomes too large (larger than 400 KB), or a local secondary
	//    index (LSI) becomes too large, or a similar validation error occurs because
	//    of changes made by the transaction.
	//
	//    * There is a user error, such as an invalid data format.
	//
	// DynamoDB rejects a TransactGetItems request under the following circumstances:
	//
	//    * There is an ongoing TransactGetItems operation that conflicts with a
	//    concurrent PutItem, UpdateItem, DeleteItem or TransactWriteItems request.
	//    In this case the TransactGetItems operation fails with a TransactionCanceledException.
	//
	//    * A table in the TransactGetItems request is in a different account or
	//    region.
	//
	//    * There is insufficient provisioned capacity for the transaction to be
	//    completed.
	//
	//    * There is a user error, such as an invalid data format.
	ErrCodeTransactionCanceledException = "TransactionCanceledException"

	// ErrCodeTransactionConflictException for service response error code
	// "TransactionConflictException".
	//
	// Operation was rejected because there is an ongoing transaction for the item.
	ErrCodeTransactionConflictException = "TransactionConflictException"

	// ErrCodeTransactionInProgressException for service response error code
	// "TransactionInProgressException".
	//
	// The transaction with the given request token is already in progress.
	ErrCodeTransactionInProgressException = "TransactionInProgressException"
)