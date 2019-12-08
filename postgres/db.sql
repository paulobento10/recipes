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
-- Name: ingredients; Type: TABLE; Schema: public; Owner: docker
--

CREATE TABLE public.ingredients (
    ingredient_id integer NOT NULL,
    ingredient_name character varying(50) NOT NULL
);


ALTER TABLE public.ingredients OWNER TO docker;

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
-- Name: recipes; Type: TABLE; Schema: public; Owner: docker
--

CREATE TABLE public.recipes (
    recipe_id integer NOT NULL,
    recipe_name character varying(50) NOT NULL,
    recipe_description character varying(500) NOT NULL,
    user_id integer NOT NULL
);


ALTER TABLE public.recipes OWNER TO docker;

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
-- Data for Name: directions; Type: TABLE DATA; Schema: public; Owner: docker
--

COPY public.directions (direction_id, direction_details, direction_order, recipe_id) FROM stdin;
\.


--
-- Data for Name: ingredients; Type: TABLE DATA; Schema: public; Owner: docker
--

COPY public.ingredients (ingredient_id, ingredient_name) FROM stdin;
\.


--
-- Data for Name: recipeingredients; Type: TABLE DATA; Schema: public; Owner: docker
--

COPY public.recipeingredients (recipeingredient_id, ingredient_id, recipe_id) FROM stdin;
\.


--
-- Data for Name: recipes; Type: TABLE DATA; Schema: public; Owner: docker
--

COPY public.recipes (recipe_id, recipe_name, recipe_description, user_id) FROM stdin;
\.


--
-- Data for Name: users; Type: TABLE DATA; Schema: public; Owner: docker
--

COPY public.users (user_id, user_name, email, password) FROM stdin;
\.


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

