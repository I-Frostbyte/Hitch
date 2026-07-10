CREATE TYPE "user_status" AS ENUM (
    'USER_STATUS_CREATED',
    'USER_STATUS_APPROVAL_REQUESTED',
    'USER_STATUS_APPROVED',
    'USER_STATUS_REJECTED'
);

CREATE TYPE "identity_document" AS ENUM (
    'IDENTITY_DOCUMENT_NATIONAL_ID',
    'IDENTITY_DOCUMENT_PASSPORT',
    'IDENTITY_DOCUMENT_DRIVER_LICENSE'
);

CREATE TABLE "users" (
    "id" UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    "status" user_status NOT NULL,
    "username" TEXT NOT NULL,
    "phone_number" TEXT NOT NULL UNIQUE,
    "password_hash" TEXT NOT NULL,
    "identity_document" identity_document NOT NULL,
    "created_at" TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    "updated_at" TIMESTAMPTZ NOT NULL DEFAULT NOW()
);