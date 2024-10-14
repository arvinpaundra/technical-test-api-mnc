--
-- PostgreSQL database dump
--

-- Dumped from database version 16.3
-- Dumped by pg_dump version 16.3

SET statement_timeout = 0;
SET lock_timeout = 0;
SET idle_in_transaction_session_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SELECT pg_catalog.set_config('search_path', '', false);
SET check_function_bodies = false;
SET xmloption = content;
SET client_min_messages = warning;
SET row_security = off;

SET default_tablespace = '';

SET default_table_access_method = heap;

--
-- Name: schema_migrations; Type: TABLE; Schema: public; Owner: root
--

CREATE TABLE public.schema_migrations (
    version bigint NOT NULL,
    dirty boolean NOT NULL
);


ALTER TABLE public.schema_migrations OWNER TO root;

--
-- Name: sessions; Type: TABLE; Schema: public; Owner: root
--

CREATE TABLE public.sessions (
    id uuid NOT NULL,
    user_id uuid NOT NULL,
    access_token text NOT NULL,
    refresh_token text,
    created_date timestamp without time zone,
    updated_date timestamp without time zone
);


ALTER TABLE public.sessions OWNER TO root;

--
-- Name: transaction_histories; Type: TABLE; Schema: public; Owner: root
--

CREATE TABLE public.transaction_histories (
    id uuid NOT NULL,
    user_id uuid NOT NULL,
    transaction_type character varying(10) NOT NULL,
    amount double precision NOT NULL,
    balance_before double precision NOT NULL,
    balance_after double precision NOT NULL,
    remarks character varying(255),
    reference_id character varying(255) NOT NULL,
    created_date timestamp without time zone,
    updated_date timestamp without time zone
);


ALTER TABLE public.transaction_histories OWNER TO root;

--
-- Name: transactions; Type: TABLE; Schema: public; Owner: root
--

CREATE TABLE public.transactions (
    id uuid NOT NULL,
    user_id uuid NOT NULL,
    target_user uuid,
    amount double precision NOT NULL,
    remarks character varying(255),
    category character varying(10) NOT NULL,
    status character varying(10) NOT NULL,
    created_date timestamp without time zone,
    updated_date timestamp without time zone
);


ALTER TABLE public.transactions OWNER TO root;

--
-- Name: users; Type: TABLE; Schema: public; Owner: root
--

CREATE TABLE public.users (
    id uuid NOT NULL,
    first_name character varying(100) NOT NULL,
    last_name character varying(100),
    phone_number character varying(15) NOT NULL,
    pin character varying(255) NOT NULL,
    balance double precision NOT NULL,
    address text,
    created_date timestamp without time zone,
    updated_date timestamp without time zone
);


ALTER TABLE public.users OWNER TO root;

--
-- Data for Name: schema_migrations; Type: TABLE DATA; Schema: public; Owner: root
--

COPY public.schema_migrations (version, dirty) FROM stdin;
20241014163434	f
\.


--
-- Data for Name: sessions; Type: TABLE DATA; Schema: public; Owner: root
--

COPY public.sessions (id, user_id, access_token, refresh_token, created_date, updated_date) FROM stdin;
01928c5b-4e25-76e3-9263-d399b9f4dc24	01928c5b-107b-7e8b-87a9-451e1958019a	eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoiMDE5MjhjNWItMTA3Yi03ZThiLTg3YTktNDUxZTE5NTgwMTlhIiwiZXhwIjoxNzI4OTc0ODQ3LCJpYXQiOjE3Mjg5MzE2NDd9.EG8hiq8WYz9liYQxoNGfa_68WZruQ5bsnEoyQ1GTcYI	eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoiMDE5MjhjNWItMTA3Yi03ZThiLTg3YTktNDUxZTE5NTgwMTlhIiwiZXhwIjoxNzI5MTkwODQ3LCJpYXQiOjE3Mjg5MzE2NDd9.R_3x8WCHWO0g67vF0tV5H1nNMpIAvJ24BznUll3L8-c	2024-10-15 01:47:27.013451	2024-10-15 01:47:27.013451
01928c5c-e270-7356-a514-3d2be57fb37a	01928c5b-38c1-7cbd-9043-75a35779ba6d	eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoiMDE5MjhjNWItMzhjMS03Y2JkLTkwNDMtNzVhMzU3NzliYTZkIiwiZXhwIjoxNzI4OTc0OTUwLCJpYXQiOjE3Mjg5MzE3NTB9.0kVIvoj6o48IBkMHc6mJz340wdzwKQs1QziFj0T6jj0	eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoiMDE5MjhjNWItMzhjMS03Y2JkLTkwNDMtNzVhMzU3NzliYTZkIiwiZXhwIjoxNzI5MTkwOTUwLCJpYXQiOjE3Mjg5MzE3NTB9.W1tgWUX16k2dAvuH7Y2rCnzGTrwH0ED89dFLMHfiqdg	2024-10-15 01:49:10.512218	2024-10-15 01:49:10.512218
\.


--
-- Data for Name: transaction_histories; Type: TABLE DATA; Schema: public; Owner: root
--

COPY public.transaction_histories (id, user_id, transaction_type, amount, balance_before, balance_after, remarks, reference_id, created_date, updated_date) FROM stdin;
01928c5b-9de5-7f39-8cfa-832005c0e246	01928c5b-107b-7e8b-87a9-451e1958019a	CREDIT	1000000	0	1000000		01928c5b-9ddb-7167-a4a9-688990cec715	2024-10-15 01:47:47.429998	2024-10-15 01:47:47.429998
01928c5b-d3a7-7ca0-8d52-9937f84c5dce	01928c5b-107b-7e8b-87a9-451e1958019a	DEBIT	100000	1000000	900000	Pulsa Telkomsel 100K	01928c5b-d3a6-739a-9a73-9bbcca4cbb1c	2024-10-15 01:48:01.191828	2024-10-15 01:48:01.191828
01928c5c-5c3f-7e1d-98c2-88a6a826d125	01928c5b-107b-7e8b-87a9-451e1958019a	DEBIT	100000	900000	800000	Hadiah ultah	01928c5c-5c38-778c-8323-e8e7a3b4e6bb	2024-10-15 01:48:36.159926	2024-10-15 01:48:36.159926
01928c5c-5c44-70e9-8794-2bdbc19d0e12	01928c5b-38c1-7cbd-9043-75a35779ba6d	CREDIT	100000	0	100000	Hadiah ultah	01928c5c-5c42-7617-a620-f393c79effce	2024-10-15 01:48:36.16406	2024-10-15 01:48:36.16406
01928c5c-92ec-769b-8045-cc48525d9355	01928c5b-107b-7e8b-87a9-451e1958019a	DEBIT	400000	800000	400000	Hadiah ultah	01928c5c-92e5-73d9-9607-23e5a691d573	2024-10-15 01:48:50.156433	2024-10-15 01:48:50.156433
01928c5c-92f0-7850-8e71-29b93b09cfbb	01928c5b-38c1-7cbd-9043-75a35779ba6d	CREDIT	400000	100000	500000	Hadiah ultah	01928c5c-92ef-720a-9b23-1f42550f6885	2024-10-15 01:48:50.160545	2024-10-15 01:48:50.160545
01928c5d-309b-7aa6-8041-731d67b77160	01928c5b-38c1-7cbd-9043-75a35779ba6d	DEBIT	50000	500000	450000	Pulsa Telkomsel 100K	01928c5d-3094-7687-9250-d0c130e916e7	2024-10-15 01:49:30.523698	2024-10-15 01:49:30.523698
\.


--
-- Data for Name: transactions; Type: TABLE DATA; Schema: public; Owner: root
--

COPY public.transactions (id, user_id, target_user, amount, remarks, category, status, created_date, updated_date) FROM stdin;
01928c5b-9ddb-7167-a4a9-688990cec715	01928c5b-107b-7e8b-87a9-451e1958019a	00000000-0000-0000-0000-000000000000	1000000		topup	PENDING	2024-10-15 01:47:47.419092	2024-10-15 01:47:47.419092
01928c5b-d3a6-739a-9a73-9bbcca4cbb1c	01928c5b-107b-7e8b-87a9-451e1958019a	00000000-0000-0000-0000-000000000000	100000	Pulsa Telkomsel 100K	payment	PENDING	2024-10-15 01:48:01.190236	2024-10-15 01:48:01.190236
01928c5c-5c38-778c-8323-e8e7a3b4e6bb	01928c5b-107b-7e8b-87a9-451e1958019a	01928c5b-38c1-7cbd-9043-75a35779ba6d	100000	Hadiah ultah	transfer	PENDING	2024-10-15 01:48:36.152494	2024-10-15 01:48:36.152494
01928c5c-5c42-7617-a620-f393c79effce	01928c5b-38c1-7cbd-9043-75a35779ba6d	00000000-0000-0000-0000-000000000000	100000	Hadiah ultah	transfer	PENDING	2024-10-15 01:48:36.162399	2024-10-15 01:48:36.162399
01928c5c-92e5-73d9-9607-23e5a691d573	01928c5b-107b-7e8b-87a9-451e1958019a	01928c5b-38c1-7cbd-9043-75a35779ba6d	400000	Hadiah ultah	transfer	PENDING	2024-10-15 01:48:50.149252	2024-10-15 01:48:50.149252
01928c5c-92ef-720a-9b23-1f42550f6885	01928c5b-38c1-7cbd-9043-75a35779ba6d	00000000-0000-0000-0000-000000000000	400000	Hadiah ultah	transfer	PENDING	2024-10-15 01:48:50.159133	2024-10-15 01:48:50.159133
01928c5d-3094-7687-9250-d0c130e916e7	01928c5b-38c1-7cbd-9043-75a35779ba6d	00000000-0000-0000-0000-000000000000	50000	Pulsa Telkomsel 100K	payment	PENDING	2024-10-15 01:49:30.516428	2024-10-15 01:49:30.516428
\.


--
-- Data for Name: users; Type: TABLE DATA; Schema: public; Owner: root
--

COPY public.users (id, first_name, last_name, phone_number, pin, balance, address, created_date, updated_date) FROM stdin;
01928c5b-107b-7e8b-87a9-451e1958019a	Sopran	Saprian	0811255555	$2a$10$hlTIE9vI0eLt1k7RCIuK4eM9wV3GrJ/kAnNCszkvbwSsIlwLPsA7C	400000	Jl. Kebon Manis No. 1	2024-10-15 01:47:11.227953	2024-10-15 01:47:11.227953
01928c5b-38c1-7cbd-9043-75a35779ba6d	Guntur	Saputro	0811255501	$2a$10$LriHzw6kPkicncX26KBkNe0Dpuy/nFNkI9ugLNZlM..J1rZhqNhJ2	450000	Jl. Kebon Sirih No. 1	2024-10-15 01:47:21.537835	2024-10-15 01:47:21.537835
\.


--
-- Name: schema_migrations schema_migrations_pkey; Type: CONSTRAINT; Schema: public; Owner: root
--

ALTER TABLE ONLY public.schema_migrations
    ADD CONSTRAINT schema_migrations_pkey PRIMARY KEY (version);


--
-- Name: sessions sessions_pkey; Type: CONSTRAINT; Schema: public; Owner: root
--

ALTER TABLE ONLY public.sessions
    ADD CONSTRAINT sessions_pkey PRIMARY KEY (id);


--
-- Name: transaction_histories transaction_histories_pkey; Type: CONSTRAINT; Schema: public; Owner: root
--

ALTER TABLE ONLY public.transaction_histories
    ADD CONSTRAINT transaction_histories_pkey PRIMARY KEY (id);


--
-- Name: transactions transactions_pkey; Type: CONSTRAINT; Schema: public; Owner: root
--

ALTER TABLE ONLY public.transactions
    ADD CONSTRAINT transactions_pkey PRIMARY KEY (id);


--
-- Name: users users_phone_number_key; Type: CONSTRAINT; Schema: public; Owner: root
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_phone_number_key UNIQUE (phone_number);


--
-- Name: users users_pkey; Type: CONSTRAINT; Schema: public; Owner: root
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_pkey PRIMARY KEY (id);


--
-- Name: sessions sessions_user_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: root
--

ALTER TABLE ONLY public.sessions
    ADD CONSTRAINT sessions_user_id_fkey FOREIGN KEY (user_id) REFERENCES public.users(id);


--
-- Name: transaction_histories transaction_histories_user_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: root
--

ALTER TABLE ONLY public.transaction_histories
    ADD CONSTRAINT transaction_histories_user_id_fkey FOREIGN KEY (user_id) REFERENCES public.users(id);


--
-- Name: transactions transactions_user_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: root
--

ALTER TABLE ONLY public.transactions
    ADD CONSTRAINT transactions_user_id_fkey FOREIGN KEY (user_id) REFERENCES public.users(id);


--
-- PostgreSQL database dump complete
--

