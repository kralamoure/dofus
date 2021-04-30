--
-- PostgreSQL database dump
--

-- Dumped from database version 13.0 (Debian 13.0-1.pgdg100+1)
-- Dumped by pg_dump version 13.2

SET
statement_timeout = 0;
SET
lock_timeout = 0;
SET
idle_in_transaction_session_timeout = 0;
SET
client_encoding = 'UTF8';
SET
standard_conforming_strings = on;
SELECT pg_catalog.set_config('search_path', '', false);
SET
check_function_bodies = false;
SET
xmloption = content;
SET
client_min_messages = warning;
SET
row_security = off;

--
-- Name: dofus; Type: SCHEMA; Schema: -; Owner: postgres
--

CREATE SCHEMA dofus;


ALTER
SCHEMA dofus OWNER TO postgres;

SET
default_tablespace = '';

SET
default_table_access_method = heap;

--
-- Name: accounts; Type: TABLE; Schema: dofus; Owner: postgres
--

CREATE TABLE dofus.accounts
(
    id           uuid                     DEFAULT extensions.uuid_generate_v4() NOT NULL,
    name         extensions.citext                                              NOT NULL,
    subscription timestamp with time zone DEFAULT '0001-01-01 00:00:00+00':: timestamp with time zone NOT NULL,
    admin        boolean                  DEFAULT false                         NOT NULL,
    user_id      uuid                                                           NOT NULL,
    last_access  timestamp with time zone DEFAULT '0001-01-01 00:00:00+00':: timestamp with time zone NOT NULL,
    last_ip      text                     DEFAULT ''::text NOT NULL
);


ALTER TABLE dofus.accounts OWNER TO postgres;

--
-- Name: users; Type: TABLE; Schema: dofus; Owner: postgres
--

CREATE TABLE dofus.users
(
    id              uuid    DEFAULT extensions.uuid_generate_v4() NOT NULL,
    email           extensions.citext                             NOT NULL,
    nickname        extensions.citext                             NOT NULL,
    hash            text                                          NOT NULL,
    secret_question text                                          NOT NULL,
    secret_answer   text                                          NOT NULL,
    gender          integer                                       NOT NULL,
    community       integer DEFAULT 2                             NOT NULL,
    chat_channels   text    DEFAULT '@i*#$p%!?:^'::text NOT NULL
);


ALTER TABLE dofus.users OWNER TO postgres;

--
-- Data for Name: accounts; Type: TABLE DATA; Schema: dofus; Owner: postgres
--

COPY dofus.accounts (id, name, subscription, admin, user_id, last_access, last_ip) FROM stdin;
42d18784
-846d-4848-9a30-1b55f68f0076	admin	2022-01-01 00:00:00+00	t	8f2a7222-46f0-473b-9ab8-240782c43cbb	2021-04-28 03:08:33.943002+00	127.0.0.1
854d7f77-813d-450f-ae4c-f752d2e356ad	amulecregg	2021-01-01 00:00:00+00	t	8f2a7222-46f0-473b-9ab8-240782c43cbb	2021-04-28 03:08:33.943002+00	127.0.0.1
\.


--
-- Data for Name: users; Type: TABLE DATA; Schema: dofus; Owner: postgres
--

COPY dofus.users (id, email, nickname, hash, secret_question, secret_answer, gender, community, chat_channels) FROM stdin;
8f2a7222
-46f0-473b-9ab8-240782c43cbb	amulecregg@gmail.com	raulg	$argon2id$v=19$m=65536,t=1,p=2$kMLaZ9ovGkPa6+42pjUfpw$u2TkS/vHV4dhqihTu/U6WV04d5y28VWua3ZFGt31hQ0	question	answer	0	4	#$pi*:!^?%
\.


--
-- Name: accounts accounts_name_key; Type: CONSTRAINT; Schema: dofus; Owner: postgres
--

ALTER TABLE ONLY dofus.accounts
    ADD CONSTRAINT accounts_name_key UNIQUE (name);


--
-- Name: accounts accounts_pkey; Type: CONSTRAINT; Schema: dofus; Owner: postgres
--

ALTER TABLE ONLY dofus.accounts
    ADD CONSTRAINT accounts_pkey PRIMARY KEY (id);


--
-- Name: users users_email_key; Type: CONSTRAINT; Schema: dofus; Owner: postgres
--

ALTER TABLE ONLY dofus.users
    ADD CONSTRAINT users_email_key UNIQUE (email);


--
-- Name: users users_nickname_key; Type: CONSTRAINT; Schema: dofus; Owner: postgres
--

ALTER TABLE ONLY dofus.users
    ADD CONSTRAINT users_nickname_key UNIQUE (nickname);


--
-- Name: users users_pkey; Type: CONSTRAINT; Schema: dofus; Owner: postgres
--

ALTER TABLE ONLY dofus.users
    ADD CONSTRAINT users_pkey PRIMARY KEY (id);


--
-- Name: accounts accounts_user_id_fkey; Type: FK CONSTRAINT; Schema: dofus; Owner: postgres
--

ALTER TABLE ONLY dofus.accounts
    ADD CONSTRAINT accounts_user_id_fkey FOREIGN KEY (user_id) REFERENCES dofus.users(id);


--
-- PostgreSQL database dump complete
--
