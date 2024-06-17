--
-- PostgreSQL database dump
--

-- Dumped from database version 15.5 (Homebrew)
-- Dumped by pg_dump version 15.5 (Homebrew)

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

--
-- Name: trigger_set_timestamp(); Type: FUNCTION; Schema: public; Owner: rudilesmana
--

CREATE FUNCTION public.trigger_set_timestamp() RETURNS trigger
    LANGUAGE plpgsql
    AS $$
BEGIN
    NEW.created_on = NOW();
    NEW.updated_on = NOW();
    RETURN NEW;
END;
$$;


ALTER FUNCTION public.trigger_set_timestamp() OWNER TO rudilesmana;

SET default_tablespace = '';

SET default_table_access_method = heap;

--
-- Name: orders; Type: TABLE; Schema: public; Owner: rudilesmana
--

CREATE TABLE public.orders (
    id integer NOT NULL,
    product_id integer NOT NULL,
    user_id integer NOT NULL,
    amount character varying(15) NOT NULL,
    total integer NOT NULL,
    created_on timestamp without time zone NOT NULL,
    updated_on timestamp without time zone NOT NULL,
    deleted_on timestamp without time zone,
    status character varying(20) NOT NULL
);


ALTER TABLE public.orders OWNER TO rudilesmana;

--
-- Name: orders_history; Type: TABLE; Schema: public; Owner: rudilesmana
--

CREATE TABLE public.orders_history (
    id integer NOT NULL,
    status character varying(15) NOT NULL,
    created_on timestamp without time zone NOT NULL,
    updated_on timestamp without time zone NOT NULL,
    deleted_on timestamp without time zone,
    collect_order character varying(255),
    user_id integer NOT NULL
);


ALTER TABLE public.orders_history OWNER TO rudilesmana;

--
-- Name: orders_history_id_seq; Type: SEQUENCE; Schema: public; Owner: rudilesmana
--

CREATE SEQUENCE public.orders_history_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.orders_history_id_seq OWNER TO rudilesmana;

--
-- Name: orders_history_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: rudilesmana
--

ALTER SEQUENCE public.orders_history_id_seq OWNED BY public.orders_history.id;


--
-- Name: orders_order_id_seq; Type: SEQUENCE; Schema: public; Owner: rudilesmana
--

CREATE SEQUENCE public.orders_order_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.orders_order_id_seq OWNER TO rudilesmana;

--
-- Name: orders_order_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: rudilesmana
--

ALTER SEQUENCE public.orders_order_id_seq OWNED BY public.orders.id;


--
-- Name: products; Type: TABLE; Schema: public; Owner: rudilesmana
--

CREATE TABLE public.products (
    id integer NOT NULL,
    name character varying(50) NOT NULL,
    description character varying(255) NOT NULL,
    amount character varying(15) NOT NULL,
    stok integer NOT NULL,
    created_on timestamp without time zone NOT NULL,
    updated_on timestamp without time zone NOT NULL,
    deleted_on timestamp without time zone
);


ALTER TABLE public.products OWNER TO rudilesmana;

--
-- Name: products_product_id_seq; Type: SEQUENCE; Schema: public; Owner: rudilesmana
--

CREATE SEQUENCE public.products_product_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.products_product_id_seq OWNER TO rudilesmana;

--
-- Name: products_product_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: rudilesmana
--

ALTER SEQUENCE public.products_product_id_seq OWNED BY public.products.id;


--
-- Name: users; Type: TABLE; Schema: public; Owner: rudilesmana
--

CREATE TABLE public.users (
    id integer NOT NULL,
    username character varying(50) NOT NULL,
    password text NOT NULL,
    email character varying(255) NOT NULL,
    created_on timestamp without time zone NOT NULL,
    updated_on timestamp without time zone NOT NULL,
    deleted_on timestamp without time zone
);


ALTER TABLE public.users OWNER TO rudilesmana;

--
-- Name: users_user_id_seq; Type: SEQUENCE; Schema: public; Owner: rudilesmana
--

CREATE SEQUENCE public.users_user_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.users_user_id_seq OWNER TO rudilesmana;

--
-- Name: users_user_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: rudilesmana
--

ALTER SEQUENCE public.users_user_id_seq OWNED BY public.users.id;


--
-- Name: orders id; Type: DEFAULT; Schema: public; Owner: rudilesmana
--

ALTER TABLE ONLY public.orders ALTER COLUMN id SET DEFAULT nextval('public.orders_order_id_seq'::regclass);


--
-- Name: orders_history id; Type: DEFAULT; Schema: public; Owner: rudilesmana
--

ALTER TABLE ONLY public.orders_history ALTER COLUMN id SET DEFAULT nextval('public.orders_history_id_seq'::regclass);


--
-- Name: products id; Type: DEFAULT; Schema: public; Owner: rudilesmana
--

ALTER TABLE ONLY public.products ALTER COLUMN id SET DEFAULT nextval('public.products_product_id_seq'::regclass);


--
-- Name: users id; Type: DEFAULT; Schema: public; Owner: rudilesmana
--

ALTER TABLE ONLY public.users ALTER COLUMN id SET DEFAULT nextval('public.users_user_id_seq'::regclass);


--
-- Data for Name: orders; Type: TABLE DATA; Schema: public; Owner: rudilesmana
--

COPY public.orders (id, product_id, user_id, amount, total, created_on, updated_on, deleted_on, status) FROM stdin;
5	2	3	2000000	1	2024-06-15 11:37:38.441352	2024-06-15 11:37:38.441352	\N	CART
7	3	3	3000000	1	2024-06-15 22:47:58.262713	2024-06-15 22:47:58.262713	\N	CART
9	2	3	2000000	1	2024-06-16 01:41:40.632435	2024-06-16 01:41:40.632435	\N	CART
11	4	3	3000000	1	2024-06-16 01:41:46.482856	2024-06-16 01:41:46.482856	\N	CART
14	8	3	30000	1	2024-06-17 22:36:31.765169	2024-06-17 22:36:31.765169	\N	CART
16	3	3	3000000	1	2024-06-17 22:47:11.466987	2024-06-17 22:47:11.466987	\N	CART
19	8	3	30000	1	2024-06-17 22:49:07.549444	2024-06-17 22:49:07.549444	\N	CART
15	7	3	30000	1	2024-06-17 22:36:33.930341	2024-06-17 22:36:33.930341	\N	PAID
2	6	3	25000	1	2024-06-15 08:05:35.402722	2024-06-15 08:05:35.402722	\N	PAID
6	6	3	25000	1	2024-06-15 11:39:00.661824	2024-06-15 11:39:00.661824	\N	PAID
12	6	3	25000	1	2024-06-16 02:14:34.031614	2024-06-16 02:14:34.031614	\N	PAID
18	6	3	25000	1	2024-06-17 22:49:05.462719	2024-06-17 22:49:05.462719	\N	PAID
17	2	3	2000000	1	2024-06-17 22:49:03.313832	2024-06-17 22:49:03.313832	\N	CART
1	1	3	1500000	1	2024-06-15 08:00:46.358179	2024-06-15 08:00:46.358179	\N	PAID
3	1	3	1500000	1	2024-06-15 08:08:43.216034	2024-06-15 08:08:43.216034	\N	PAID
4	1	3	1500000	1	2024-06-15 11:37:18.340501	2024-06-15 11:37:18.340501	\N	PAID
8	1	3	1500000	1	2024-06-15 23:18:28.345083	2024-06-15 23:18:28.345083	\N	PAID
10	1	3	1500000	1	2024-06-16 01:41:42.913841	2024-06-16 01:41:42.913841	\N	PAID
13	1	3	1500000	1	2024-06-16 02:14:36.28103	2024-06-16 02:14:36.28103	\N	PAID
\.


--
-- Data for Name: orders_history; Type: TABLE DATA; Schema: public; Owner: rudilesmana
--

COPY public.orders_history (id, status, created_on, updated_on, deleted_on, collect_order, user_id) FROM stdin;
31	PAID	2024-06-16 23:50:10.965	2024-06-16 23:50:10.965	\N	15,2,6,12,18	3
32	PAID	2024-06-17 23:51:24.693322	2024-06-17 23:51:24.693322	\N	1,3,4,8,10,13	3
\.


--
-- Data for Name: products; Type: TABLE DATA; Schema: public; Owner: rudilesmana
--

COPY public.products (id, name, description, amount, stok, created_on, updated_on, deleted_on) FROM stdin;
1	Lining Raket	original	1500000	44	2023-08-08 03:12:24.37	2023-08-08 03:12:24.37	\N
11	Sample Product	Sample Description	1000	50	2023-08-11 00:15:43.055895	2023-08-11 00:15:43.055895	2023-08-15 23:06:03.990684
2	Mizuno Sepatu	original	2000000	50	2023-08-08 03:12:24.37	2023-08-08 03:12:24.37	\N
3	Nike Sepatu	original	3000000	50	2023-08-08 03:12:24.37	2023-08-08 03:12:24.37	\N
4	Adidas Sepatu	original	3000000	50	2023-08-08 03:12:24.37	2023-08-08 03:12:24.37	\N
8	RS Celana	premium	30000	50	2023-08-08 03:12:24.37	2023-08-08 03:12:24.37	\N
7	Yonex Kaos	premium	30000	49	2023-08-08 03:12:24.37	2023-08-08 03:12:24.37	\N
6	Victor Celana	premium	25000	46	2023-08-08 03:12:24.37	2023-08-08 03:12:24.37	\N
5	Lining Kaos	premium	15000	50	2023-08-08 03:12:24.37	2023-08-08 03:12:24.37	\N
\.


--
-- Data for Name: users; Type: TABLE DATA; Schema: public; Owner: rudilesmana
--

COPY public.users (id, username, password, email, created_on, updated_on, deleted_on) FROM stdin;
3	Rudi	$2a$12$cd/uxz/udUUwYY/Cy3Z1quOk1sdaXmGfoTCJ/oKrX/r9TbfqYx1U2	rudi@gmail.com	2023-08-08 03:12:24.37	2023-08-08 03:12:24.37	\N
4	Anwar	$2a$12$cd/uxz/udUUwYY/Cy3Z1quOk1sdaXmGfoTCJ/oKrX/r9TbfqYx1U2	anwar@gmail.com	2023-08-08 03:12:24.37	2023-08-08 03:12:24.37	\N
5	Roni	$2a$12$cd/uxz/udUUwYY/Cy3Z1quOk1sdaXmGfoTCJ/oKrX/r9TbfqYx1U2	roni@gmail.com	2023-08-08 03:12:24.37	2023-08-08 03:12:24.37	\N
6	Firza	$2a$12$cd/uxz/udUUwYY/Cy3Z1quOk1sdaXmGfoTCJ/oKrX/r9TbfqYx1U2	firza@gmail.com	2023-08-08 03:12:24.37	2023-08-08 03:12:24.37	\N
\.


--
-- Name: orders_history_id_seq; Type: SEQUENCE SET; Schema: public; Owner: rudilesmana
--

SELECT pg_catalog.setval('public.orders_history_id_seq', 32, true);


--
-- Name: orders_order_id_seq; Type: SEQUENCE SET; Schema: public; Owner: rudilesmana
--

SELECT pg_catalog.setval('public.orders_order_id_seq', 19, true);


--
-- Name: products_product_id_seq; Type: SEQUENCE SET; Schema: public; Owner: rudilesmana
--

SELECT pg_catalog.setval('public.products_product_id_seq', 26, true);


--
-- Name: users_user_id_seq; Type: SEQUENCE SET; Schema: public; Owner: rudilesmana
--

SELECT pg_catalog.setval('public.users_user_id_seq', 6, true);


--
-- Name: orders_history orders_history_pkey; Type: CONSTRAINT; Schema: public; Owner: rudilesmana
--

ALTER TABLE ONLY public.orders_history
    ADD CONSTRAINT orders_history_pkey PRIMARY KEY (id);


--
-- Name: orders orders_pkey; Type: CONSTRAINT; Schema: public; Owner: rudilesmana
--

ALTER TABLE ONLY public.orders
    ADD CONSTRAINT orders_pkey PRIMARY KEY (id);


--
-- Name: products products_pkey; Type: CONSTRAINT; Schema: public; Owner: rudilesmana
--

ALTER TABLE ONLY public.products
    ADD CONSTRAINT products_pkey PRIMARY KEY (id);


--
-- Name: users users_pkey; Type: CONSTRAINT; Schema: public; Owner: rudilesmana
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_pkey PRIMARY KEY (id);


--
-- Name: orders set_timestamp; Type: TRIGGER; Schema: public; Owner: rudilesmana
--

CREATE TRIGGER set_timestamp BEFORE INSERT ON public.orders FOR EACH ROW EXECUTE FUNCTION public.trigger_set_timestamp();


--
-- Name: orders_history set_timestamp; Type: TRIGGER; Schema: public; Owner: rudilesmana
--

CREATE TRIGGER set_timestamp BEFORE INSERT ON public.orders_history FOR EACH ROW EXECUTE FUNCTION public.trigger_set_timestamp();


--
-- Name: products set_timestamp; Type: TRIGGER; Schema: public; Owner: rudilesmana
--

CREATE TRIGGER set_timestamp BEFORE INSERT ON public.products FOR EACH ROW EXECUTE FUNCTION public.trigger_set_timestamp();


--
-- Name: users set_timestamp; Type: TRIGGER; Schema: public; Owner: rudilesmana
--

CREATE TRIGGER set_timestamp BEFORE INSERT ON public.users FOR EACH ROW EXECUTE FUNCTION public.trigger_set_timestamp();


--
-- PostgreSQL database dump complete
--

