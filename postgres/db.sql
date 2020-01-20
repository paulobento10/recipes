--
-- PostgreSQL database dump
--

-- Dumped from database version 11.5
-- Dumped by pg_dump version 11.5

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

/*ALTER TABLE ONLY public.recipes DROP CONSTRAINT recipes_user_id_fkey;
ALTER TABLE ONLY public.recipeingredients DROP CONSTRAINT recipeingredients_recipe_id_fkey;
ALTER TABLE ONLY public.recipeingredients DROP CONSTRAINT recipeingredients_ingredient_id_fkey;
ALTER TABLE ONLY public.ingredients DROP CONSTRAINT ingredient_user_id_fkey;
ALTER TABLE ONLY public.directions DROP CONSTRAINT directions_recipe_id_fkey;
ALTER TABLE ONLY public.users DROP CONSTRAINT users_user_name_key;
ALTER TABLE ONLY public.users DROP CONSTRAINT users_pkey;
ALTER TABLE ONLY public.recipes DROP CONSTRAINT recipes_recipe_name_key;
ALTER TABLE ONLY public.recipes DROP CONSTRAINT recipes_pkey;
ALTER TABLE ONLY public.recipeingredients DROP CONSTRAINT recipeingredients_pkey;
ALTER TABLE ONLY public.ingredients DROP CONSTRAINT ingredients_pkey;
ALTER TABLE ONLY public.ingredients DROP CONSTRAINT ingredients_ingredient_name_key;
ALTER TABLE ONLY public.directions DROP CONSTRAINT directions_pkey;
ALTER TABLE public.users ALTER COLUMN user_id DROP DEFAULT;
ALTER TABLE public.recipes ALTER COLUMN recipe_id DROP DEFAULT;
ALTER TABLE public.recipeingredients ALTER COLUMN recipeingredient_id DROP DEFAULT;
ALTER TABLE public.ingredients ALTER COLUMN ingredient_id DROP DEFAULT;
ALTER TABLE public.directions ALTER COLUMN direction_id DROP DEFAULT;
DROP SEQUENCE public.users_user_id_seq;
DROP TABLE public.users;
DROP SEQUENCE public.recipes_recipe_id_seq;
DROP TABLE public.recipes;
DROP SEQUENCE public.recipeingredients_recipeingredient_id_seq;
DROP TABLE public.recipeingredients;
DROP SEQUENCE public.ingredients_ingredient_id_seq;
DROP TABLE public.ingredients;
DROP SEQUENCE public.directions_direction_id_seq;
DROP TABLE public.directions;*/
SET default_tablespace = '';

SET default_with_oids = false;

--
-- Name: directions; Type: TABLE; Schema: public; Owner: docker
--

CREATE TABLE public.directions (
    direction_id integer NOT NULL,
    direction_details character varying(500) NOT NULL,
    direction_order integer NOT NULL,
    recipe_id integer NOT NULL
);


ALTER TABLE public.directions OWNER TO docker;

--
-- Name: directions_direction_id_seq; Type: SEQUENCE; Schema: public; Owner: docker
--

CREATE SEQUENCE public.directions_direction_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.directions_direction_id_seq OWNER TO docker;

--
-- Name: directions_direction_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: docker
--

ALTER SEQUENCE public.directions_direction_id_seq OWNED BY public.directions.direction_id;


--
-- Name: ingredients; Type: TABLE; Schema: public; Owner: docker
--

CREATE TABLE public.ingredients (
    ingredient_id integer NOT NULL,
    ingredient_name character varying(50) NOT NULL,
    kcal text,
    user_id integer
);


ALTER TABLE public.ingredients OWNER TO docker;

--
-- Name: ingredients_ingredient_id_seq; Type: SEQUENCE; Schema: public; Owner: docker
--

CREATE SEQUENCE public.ingredients_ingredient_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.ingredients_ingredient_id_seq OWNER TO docker;

--
-- Name: ingredients_ingredient_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: docker
--

ALTER SEQUENCE public.ingredients_ingredient_id_seq OWNED BY public.ingredients.ingredient_id;


--
-- Name: recipeingredients; Type: TABLE; Schema: public; Owner: docker
--

CREATE TABLE public.recipeingredients (
    recipeingredient_id integer NOT NULL,
    ingredient_id integer NOT NULL,
    recipe_id integer NOT NULL
);


ALTER TABLE public.recipeingredients OWNER TO docker;

--
-- Name: recipeingredients_recipeingredient_id_seq; Type: SEQUENCE; Schema: public; Owner: docker
--

CREATE SEQUENCE public.recipeingredients_recipeingredient_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.recipeingredients_recipeingredient_id_seq OWNER TO docker;

--
-- Name: recipeingredients_recipeingredient_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: docker
--

ALTER SEQUENCE public.recipeingredients_recipeingredient_id_seq OWNED BY public.recipeingredients.recipeingredient_id;


--
-- Name: recipes; Type: TABLE; Schema: public; Owner: docker
--

CREATE TABLE public.recipes (
    recipe_id integer NOT NULL,
    recipe_name character varying(50) NOT NULL,
    recipe_description character varying(500) NOT NULL,
    user_id integer NOT NULL,
    duration text,
    picture text,
    category text,
    kcal text
);


ALTER TABLE public.recipes OWNER TO docker;

--
-- Name: recipes_recipe_id_seq; Type: SEQUENCE; Schema: public; Owner: docker
--

CREATE SEQUENCE public.recipes_recipe_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.recipes_recipe_id_seq OWNER TO docker;

--
-- Name: recipes_recipe_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: docker
--

ALTER SEQUENCE public.recipes_recipe_id_seq OWNED BY public.recipes.recipe_id;


--
-- Name: users; Type: TABLE; Schema: public; Owner: docker
--

CREATE TABLE public.users (
    user_id integer NOT NULL,
    user_name character varying(50) NOT NULL,
    email character varying(50) NOT NULL,
    password character varying NOT NULL
);


ALTER TABLE public.users OWNER TO docker;

--
-- Name: users_user_id_seq; Type: SEQUENCE; Schema: public; Owner: docker
--

CREATE SEQUENCE public.users_user_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.users_user_id_seq OWNER TO docker;

--
-- Name: users_user_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: docker
--

ALTER SEQUENCE public.users_user_id_seq OWNED BY public.users.user_id;


--
-- Name: directions direction_id; Type: DEFAULT; Schema: public; Owner: docker
--

ALTER TABLE ONLY public.directions ALTER COLUMN direction_id SET DEFAULT nextval('public.directions_direction_id_seq'::regclass);


--
-- Name: ingredients ingredient_id; Type: DEFAULT; Schema: public; Owner: docker
--

ALTER TABLE ONLY public.ingredients ALTER COLUMN ingredient_id SET DEFAULT nextval('public.ingredients_ingredient_id_seq'::regclass);


--
-- Name: recipeingredients recipeingredient_id; Type: DEFAULT; Schema: public; Owner: docker
--

ALTER TABLE ONLY public.recipeingredients ALTER COLUMN recipeingredient_id SET DEFAULT nextval('public.recipeingredients_recipeingredient_id_seq'::regclass);


--
-- Name: recipes recipe_id; Type: DEFAULT; Schema: public; Owner: docker
--

ALTER TABLE ONLY public.recipes ALTER COLUMN recipe_id SET DEFAULT nextval('public.recipes_recipe_id_seq'::regclass);


--
-- Name: users user_id; Type: DEFAULT; Schema: public; Owner: docker
--

ALTER TABLE ONLY public.users ALTER COLUMN user_id SET DEFAULT nextval('public.users_user_id_seq'::regclass);


--
-- Data for Name: directions; Type: TABLE DATA; Schema: public; Owner: docker
--

COPY public.directions (direction_id, direction_details, direction_order, recipe_id) FROM stdin;
\.


--
-- Data for Name: ingredients; Type: TABLE DATA; Schema: public; Owner: docker
--

COPY public.ingredients (ingredient_id, ingredient_name, kcal, user_id) FROM stdin;
1	banana	100	1
\.


--
-- Data for Name: recipeingredients; Type: TABLE DATA; Schema: public; Owner: docker
--

COPY public.recipeingredients (recipeingredient_id, ingredient_id, recipe_id) FROM stdin;
\.


--
-- Data for Name: recipes; Type: TABLE DATA; Schema: public; Owner: docker
--

COPY public.recipes (recipe_id, recipe_name, recipe_description, user_id, duration, picture, category, kcal) FROM stdin;
\.


--
-- Data for Name: users; Type: TABLE DATA; Schema: public; Owner: docker
--

COPY public.users (user_id, user_name, email, password) FROM stdin;
1	paulo	paulo	123
\.


--
-- Name: directions_direction_id_seq; Type: SEQUENCE SET; Schema: public; Owner: docker
--

SELECT pg_catalog.setval('public.directions_direction_id_seq', 1, false);


--
-- Name: ingredients_ingredient_id_seq; Type: SEQUENCE SET; Schema: public; Owner: docker
--

SELECT pg_catalog.setval('public.ingredients_ingredient_id_seq', 3, true);


--
-- Name: recipeingredients_recipeingredient_id_seq; Type: SEQUENCE SET; Schema: public; Owner: docker
--

SELECT pg_catalog.setval('public.recipeingredients_recipeingredient_id_seq', 1, false);


--
-- Name: recipes_recipe_id_seq; Type: SEQUENCE SET; Schema: public; Owner: docker
--

SELECT pg_catalog.setval('public.recipes_recipe_id_seq', 1, false);


--
-- Name: users_user_id_seq; Type: SEQUENCE SET; Schema: public; Owner: docker
--

SELECT pg_catalog.setval('public.users_user_id_seq', 1, true);


--
-- Name: directions directions_pkey; Type: CONSTRAINT; Schema: public; Owner: docker
--

ALTER TABLE ONLY public.directions
    ADD CONSTRAINT directions_pkey PRIMARY KEY (direction_id);


--
-- Name: ingredients ingredients_ingredient_name_key; Type: CONSTRAINT; Schema: public; Owner: docker
--

ALTER TABLE ONLY public.ingredients
    ADD CONSTRAINT ingredients_ingredient_name_key UNIQUE (ingredient_name);


--
-- Name: ingredients ingredients_pkey; Type: CONSTRAINT; Schema: public; Owner: docker
--

ALTER TABLE ONLY public.ingredients
    ADD CONSTRAINT ingredients_pkey PRIMARY KEY (ingredient_id);


--
-- Name: recipeingredients recipeingredients_pkey; Type: CONSTRAINT; Schema: public; Owner: docker
--

ALTER TABLE ONLY public.recipeingredients
    ADD CONSTRAINT recipeingredients_pkey PRIMARY KEY (recipeingredient_id);


--
-- Name: recipes recipes_pkey; Type: CONSTRAINT; Schema: public; Owner: docker
--

ALTER TABLE ONLY public.recipes
    ADD CONSTRAINT recipes_pkey PRIMARY KEY (recipe_id);


--
-- Name: recipes recipes_recipe_name_key; Type: CONSTRAINT; Schema: public; Owner: docker
--

ALTER TABLE ONLY public.recipes
    ADD CONSTRAINT recipes_recipe_name_key UNIQUE (recipe_name);


--
-- Name: users users_pkey; Type: CONSTRAINT; Schema: public; Owner: docker
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_pkey PRIMARY KEY (user_id);


--
-- Name: users users_user_name_key; Type: CONSTRAINT; Schema: public; Owner: docker
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_user_name_key UNIQUE (user_name);


--
-- Name: directions directions_recipe_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: docker
--

ALTER TABLE ONLY public.directions
    ADD CONSTRAINT directions_recipe_id_fkey FOREIGN KEY (recipe_id) REFERENCES public.recipes(recipe_id);


--
-- Name: ingredients ingredient_user_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: docker
--

ALTER TABLE ONLY public.ingredients
    ADD CONSTRAINT ingredient_user_id_fkey FOREIGN KEY (user_id) REFERENCES public.users(user_id);


--
-- Name: recipeingredients recipeingredients_ingredient_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: docker
--

ALTER TABLE ONLY public.recipeingredients
    ADD CONSTRAINT recipeingredients_ingredient_id_fkey FOREIGN KEY (ingredient_id) REFERENCES public.ingredients(ingredient_id);


--
-- Name: recipeingredients recipeingredients_recipe_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: docker
--

ALTER TABLE ONLY public.recipeingredients
    ADD CONSTRAINT recipeingredients_recipe_id_fkey FOREIGN KEY (recipe_id) REFERENCES public.recipes(recipe_id);


--
-- Name: recipes recipes_user_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: docker
--

ALTER TABLE ONLY public.recipes
    ADD CONSTRAINT recipes_user_id_fkey FOREIGN KEY (user_id) REFERENCES public.users(user_id);


--
-- PostgreSQL database dump complete
--
