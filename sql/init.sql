-- ----------------------------
-- Table structure for ab_album
-- ----------------------------
DROP TABLE IF EXISTS "public"."ab_album";
CREATE TABLE "public"."ab_album" (
"id" serial NOT NULL,
"title" varchar COLLATE "default",
"artist" varchar COLLATE "default",
"price" float ,
"create_time" timestamp(6) DEFAULT (now()),
"user_id" int4
)
WITH (OIDS=FALSE)

;