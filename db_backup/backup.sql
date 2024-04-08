--
-- PostgreSQL database dump
--

-- Dumped from database version 14.3
-- Dumped by pg_dump version 14.3

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
-- Name: customer_vehicles; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.customer_vehicles (
    id bigint NOT NULL,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone,
    customer_id bigint,
    vehicle_id bigint,
    link text
);


ALTER TABLE public.customer_vehicles OWNER TO postgres;

--
-- Name: customer_vehicles_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.customer_vehicles_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.customer_vehicles_id_seq OWNER TO postgres;

--
-- Name: customer_vehicles_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.customer_vehicles_id_seq OWNED BY public.customer_vehicles.id;


--
-- Name: customers; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.customers (
    id bigint NOT NULL,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone,
    name text,
    cnpj text
);


ALTER TABLE public.customers OWNER TO postgres;

--
-- Name: customers_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.customers_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.customers_id_seq OWNER TO postgres;

--
-- Name: customers_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.customers_id_seq OWNED BY public.customers.id;


--
-- Name: vehicles; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.vehicles (
    id bigint NOT NULL,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    deleted_at timestamp with time zone,
    license_plate text,
    color text
);


ALTER TABLE public.vehicles OWNER TO postgres;

--
-- Name: vehicles_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.vehicles_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.vehicles_id_seq OWNER TO postgres;

--
-- Name: vehicles_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.vehicles_id_seq OWNED BY public.vehicles.id;


--
-- Name: customer_vehicles id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.customer_vehicles ALTER COLUMN id SET DEFAULT nextval('public.customer_vehicles_id_seq'::regclass);


--
-- Name: customers id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.customers ALTER COLUMN id SET DEFAULT nextval('public.customers_id_seq'::regclass);


--
-- Name: vehicles id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.vehicles ALTER COLUMN id SET DEFAULT nextval('public.vehicles_id_seq'::regclass);


--
-- Data for Name: customer_vehicles; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.customer_vehicles (id, created_at, updated_at, deleted_at, customer_id, vehicle_id, link) FROM stdin;
1	2024-04-07 21:22:48.033536-03	2024-04-07 21:22:48.033536-03	2024-04-07 21:26:30.031462-03	1	1	AUTONOMO
2	2024-04-07 21:29:47.253576-03	2024-04-07 21:29:47.253576-03	2024-04-07 21:32:06.534527-03	2	2	FROTA
8	2024-04-07 21:45:06.379765-03	2024-04-07 21:45:06.379765-03	2024-04-07 21:45:51.717045-03	4	1	AGREGADO
9	2024-04-07 21:45:18.617575-03	2024-04-07 21:45:18.617575-03	2024-04-07 21:45:51.717045-03	4	4	AGREGADO
10	2024-04-07 21:45:21.307922-03	2024-04-07 21:45:21.307922-03	2024-04-07 21:45:51.717045-03	4	5	AGREGADO
11	2024-04-07 21:45:25.815767-03	2024-04-07 21:45:25.815767-03	2024-04-07 21:45:51.717045-03	4	6	AGREGADO
12	2024-04-07 21:48:02.713025-03	2024-04-07 21:48:02.713025-03	2024-04-07 21:50:20.750552-03	5	7	AUTONOMO
13	2024-04-07 21:48:06.813375-03	2024-04-07 21:48:06.813375-03	2024-04-07 21:50:20.750552-03	5	8	AUTONOMO
14	2024-04-07 21:48:11.031544-03	2024-04-07 21:48:11.031544-03	2024-04-07 21:50:20.750552-03	5	9	AUTONOMO
15	2024-04-07 21:48:15.478587-03	2024-04-07 21:48:15.478587-03	2024-04-07 21:50:20.750552-03	5	10	AUTONOMO
16	2024-04-07 22:01:23.547272-03	2024-04-07 22:01:23.547272-03	\N	4	6	AUTONOMO
18	2024-04-08 08:13:40.388514-03	2024-04-08 08:13:40.388514-03	\N	2	3	AUTONOMO
3	2024-04-07 21:32:06.582846-03	2024-04-07 21:32:06.582846-03	2024-04-08 08:14:15.209649-03	3	2	FROTA
19	2024-04-08 08:14:15.255425-03	2024-04-08 08:14:15.255425-03	\N	4	2	FROTA
20	2024-04-08 08:20:46.177032-03	2024-04-08 08:20:46.177032-03	\N	3	5	AUTONOMO
22	2024-04-08 08:20:54.799-03	2024-04-08 08:20:54.799-03	\N	3	7	AUTONOMO
21	2024-04-08 08:20:51.338413-03	2024-04-08 08:20:51.338413-03	2024-04-08 08:21:07.48004-03	3	6	AUTONOMO
17	2024-04-07 22:01:42.627488-03	2024-04-07 22:01:42.627488-03	2024-04-08 08:22:25.709404-03	5	6	AUTONOMO
23	2024-04-08 08:21:41.746105-03	2024-04-08 08:21:41.746105-03	2024-04-08 08:22:25.709404-03	5	3	AGREGADO
24	2024-04-08 08:21:45.12616-03	2024-04-08 08:21:45.12616-03	2024-04-08 08:22:25.709404-03	5	4	AGREGADO
\.


--
-- Data for Name: customers; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.customers (id, created_at, updated_at, deleted_at, name, cnpj) FROM stdin;
1	2024-04-07 21:15:18.42053-03	2024-04-07 21:15:18.42053-03	\N	costumer1	1234567890
2	2024-04-07 21:15:27.936496-03	2024-04-07 21:15:27.936496-03	\N	costumer2	1234567890
3	2024-04-07 21:15:33.507124-03	2024-04-07 21:15:33.507124-03	\N	costumer3	1234567890
4	2024-04-07 21:15:37.620011-03	2024-04-07 21:15:37.620011-03	\N	costumer4	1234567890
5	2024-04-07 21:15:42.26684-03	2024-04-07 21:15:42.26684-03	\N	costumer5	1234567890
\.


--
-- Data for Name: vehicles; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.vehicles (id, created_at, updated_at, deleted_at, license_plate, color) FROM stdin;
2	2024-04-07 21:18:34.019374-03	2024-04-07 21:18:34.019374-03	\N	ABC0002	color
3	2024-04-07 21:18:38.008082-03	2024-04-07 21:18:38.008082-03	\N	ABC0003	color
4	2024-04-07 21:18:41.302765-03	2024-04-07 21:18:41.302765-03	\N	ABC0004	color
5	2024-04-07 21:18:45.919083-03	2024-04-07 21:18:45.919083-03	\N	ABC0005	color
6	2024-04-07 21:18:49.149889-03	2024-04-07 21:18:49.149889-03	\N	ABC0006	color
7	2024-04-07 21:18:54.422936-03	2024-04-07 21:18:54.422936-03	\N	ABC0007	color
8	2024-04-07 21:18:59.792442-03	2024-04-07 21:18:59.792442-03	\N	ABC0008	color
9	2024-04-07 21:19:05.357485-03	2024-04-07 21:19:05.357485-03	\N	ABC0009	color
10	2024-04-07 21:19:17.823556-03	2024-04-07 21:19:17.823556-03	\N	ABC0010	color
1	2024-04-07 21:18:30.043138-03	2024-04-08 08:09:33.932044-03	\N	ABCD0001	color
\.


--
-- Name: customer_vehicles_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.customer_vehicles_id_seq', 24, true);


--
-- Name: customers_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.customers_id_seq', 5, true);


--
-- Name: vehicles_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.vehicles_id_seq', 10, true);


--
-- Name: customer_vehicles customer_vehicles_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.customer_vehicles
    ADD CONSTRAINT customer_vehicles_pkey PRIMARY KEY (id);


--
-- Name: customers customers_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.customers
    ADD CONSTRAINT customers_pkey PRIMARY KEY (id);


--
-- Name: vehicles vehicles_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.vehicles
    ADD CONSTRAINT vehicles_pkey PRIMARY KEY (id);


--
-- Name: idx_customer_vehicles_deleted_at; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_customer_vehicles_deleted_at ON public.customer_vehicles USING btree (deleted_at);


--
-- Name: idx_customers_deleted_at; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_customers_deleted_at ON public.customers USING btree (deleted_at);


--
-- Name: idx_vehicles_deleted_at; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_vehicles_deleted_at ON public.vehicles USING btree (deleted_at);


--
-- PostgreSQL database dump complete
--

