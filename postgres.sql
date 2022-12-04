-- Adminer 4.8.1 PostgreSQL 15.1 (Debian 15.1-1.pgdg110+1) dump

DROP TABLE IF EXISTS "item";
DROP SEQUENCE IF EXISTS item_id_seq;
CREATE SEQUENCE item_id_seq INCREMENT 1 MINVALUE 1 MAXVALUE 2147483647 CACHE 1;

CREATE TABLE "public"."item" (
    "id" integer DEFAULT nextval('item_id_seq') NOT NULL,
    "name" character varying(500) NOT NULL,
    "created_at" timestamp,
    CONSTRAINT "item_name" UNIQUE ("name"),
    CONSTRAINT "item_pkey" PRIMARY KEY ("id")
) WITH (oids = false);

INSERT INTO "item" ("id", "name", "created_at") VALUES
(1,	'item_1',	'2022-11-13 15:13:10.737568');

DROP TABLE IF EXISTS "sale";
DROP SEQUENCE IF EXISTS sale_id_seq;
CREATE SEQUENCE sale_id_seq INCREMENT 1 MINVALUE 1 MAXVALUE 2147483647 CACHE 1;

CREATE TABLE "public"."sale" (
    "id" integer DEFAULT nextval('sale_id_seq') NOT NULL,
    "user_id" integer NOT NULL,
    "price" money NOT NULL,
    "created_at" timestamp,
    CONSTRAINT "sale_pkey" PRIMARY KEY ("id")
) WITH (oids = false);


DROP TABLE IF EXISTS "sale_item";
DROP SEQUENCE IF EXISTS sale_item_id_seq;
CREATE SEQUENCE sale_item_id_seq INCREMENT 1 MINVALUE 1 MAXVALUE 2147483647 CACHE 1;

CREATE TABLE "public"."sale_item" (
    "id" integer DEFAULT nextval('sale_item_id_seq') NOT NULL,
    "sale_id" integer NOT NULL,
    "item_id" integer NOT NULL,
    "amount" integer NOT NULL,
    CONSTRAINT "sale_item_pkey" PRIMARY KEY ("id")
) WITH (oids = false);


DROP TABLE IF EXISTS "user";
DROP SEQUENCE IF EXISTS user_id_seq;
CREATE SEQUENCE user_id_seq INCREMENT 1 MINVALUE 1 MAXVALUE 2147483647 CACHE 1;

CREATE TABLE "public"."user" (
    "id" integer DEFAULT nextval('user_id_seq') NOT NULL,
    "username" character varying(50) NOT NULL,
    "password" character varying(100) NOT NULL,
    "created_at" timestamp,
    CONSTRAINT "user_pkey" PRIMARY KEY ("id"),
    CONSTRAINT "user_username" UNIQUE ("username")
) WITH (oids = false);

INSERT INTO "user" ("id", "username", "password", "created_at") VALUES
(1,	'test_user',	'0b1baf69f962f0a838a90b8395dccb4d1db6fb001f033e6495dab6ab585a107e',	'2022-11-13 15:15:23.522369');

ALTER TABLE ONLY "public"."sale" ADD CONSTRAINT "sale_user_id_fkey" FOREIGN KEY (user_id) REFERENCES "user"(id) ON UPDATE CASCADE ON DELETE SET NULL NOT DEFERRABLE;

ALTER TABLE ONLY "public"."sale_item" ADD CONSTRAINT "sale_item_item_id_fkey" FOREIGN KEY (item_id) REFERENCES item(id) ON UPDATE CASCADE ON DELETE SET NULL NOT DEFERRABLE;
ALTER TABLE ONLY "public"."sale_item" ADD CONSTRAINT "sale_item_sale_id_fkey" FOREIGN KEY (sale_id) REFERENCES sale(id) ON UPDATE CASCADE ON DELETE SET NULL NOT DEFERRABLE;

-- 2022-11-19 09:27:32.245085+07
