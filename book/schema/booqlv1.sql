--
-- PostgreSQL database dump
--

-- Dumped from database version 12.1
-- Dumped by pg_dump version 12.1

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
-- Name: author; Type: TABLE; Schema: public; Owner: ferdinandusrichard
--

CREATE TABLE public.author (
    id bigint NOT NULL,
    name character varying(255) NOT NULL
);


ALTER TABLE public.author OWNER TO ferdinandusrichard;

--
-- Name: author_id_seq; Type: SEQUENCE; Schema: public; Owner: ferdinandusrichard
--

CREATE SEQUENCE public.author_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.author_id_seq OWNER TO ferdinandusrichard;

--
-- Name: author_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: ferdinandusrichard
--

ALTER SEQUENCE public.author_id_seq OWNED BY public.author.id;


--
-- Name: book; Type: TABLE; Schema: public; Owner: ferdinandusrichard
--

CREATE TABLE public.book (
    id bigint NOT NULL,
    title character varying(255) NOT NULL,
    author_id bigint NOT NULL,
    year integer NOT NULL
);


ALTER TABLE public.book OWNER TO ferdinandusrichard;

--
-- Name: book_id_seq; Type: SEQUENCE; Schema: public; Owner: ferdinandusrichard
--

CREATE SEQUENCE public.book_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.book_id_seq OWNER TO ferdinandusrichard;

--
-- Name: book_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: ferdinandusrichard
--

ALTER SEQUENCE public.book_id_seq OWNED BY public.book.id;


--
-- Name: author id; Type: DEFAULT; Schema: public; Owner: ferdinandusrichard
--

ALTER TABLE ONLY public.author ALTER COLUMN id SET DEFAULT nextval('public.author_id_seq'::regclass);


--
-- Name: book id; Type: DEFAULT; Schema: public; Owner: ferdinandusrichard
--

ALTER TABLE ONLY public.book ALTER COLUMN id SET DEFAULT nextval('public.book_id_seq'::regclass);


--
-- Data for Name: author; Type: TABLE DATA; Schema: public; Owner: ferdinandusrichard
--

COPY public.author (id, name) FROM stdin;
\.


--
-- Data for Name: book; Type: TABLE DATA; Schema: public; Owner: ferdinandusrichard
--

COPY public.book (id, title, author_id, year) FROM stdin;
\.


--
-- Name: author_id_seq; Type: SEQUENCE SET; Schema: public; Owner: ferdinandusrichard
--

SELECT pg_catalog.setval('public.author_id_seq', 1, false);


--
-- Name: book_id_seq; Type: SEQUENCE SET; Schema: public; Owner: ferdinandusrichard
--

SELECT pg_catalog.setval('public.book_id_seq', 1, false);


--
-- Name: author author_pkey; Type: CONSTRAINT; Schema: public; Owner: ferdinandusrichard
--

ALTER TABLE ONLY public.author
    ADD CONSTRAINT author_pkey PRIMARY KEY (id);


--
-- Name: book book_pkey; Type: CONSTRAINT; Schema: public; Owner: ferdinandusrichard
--

ALTER TABLE ONLY public.book
    ADD CONSTRAINT book_pkey PRIMARY KEY (id);


--
-- Name: book book_author_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: ferdinandusrichard
--

ALTER TABLE ONLY public.book
    ADD CONSTRAINT book_author_id_fkey FOREIGN KEY (author_id) REFERENCES public.author(id);


--
-- PostgreSQL database dump complete
--

