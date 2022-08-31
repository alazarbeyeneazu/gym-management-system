CREATE TABLE "users" (
  "id" serial PRIMARY KEY,
  "first_name" VARCHAR NOT NULL,
  "last_name" VARCHAR NOT NULL,
  "phone_number" VARCHAR NOT NULL,
  "email" VARCHAR NOT NULL,
  "password" VARCHAR NOT NULL,
  "created_at" VARCHAR NOT NULL,
  "state" int64 NOT NULL
);
CREATE TABLE "admins" (
  "id" serial PRIMARY KEY,
  "admin_account_id" int64 NOT NULL REFERENCES users(id)
  
);
CREATE TABLE "vehicle_types" (
  "id" serial PRIMARY KEY,
  "vehicle_name" varchar NOT NULL
);

CREATE TABLE "vehicles" (
  "id" serial PRIMARY KEY,
  "driver_id" int64 NOT NULL REFERENCES users(id),
  "vehicle_type" int64 NOT NULL REFERENCES vehicle_types(id)
);
CREATE TABLE "vehicle_service_types" (
  "id" serial PRIMARY KEY,
  "vehicle_id" int64 NOT NULL REFERENCES vehicles(id),
  "service_name" varchar NOT NULL
);


