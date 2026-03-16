

CREATE FUNCTION public.update_updated_at_column() RETURNS trigger
    LANGUAGE plpgsql
    AS $$
BEGIN
    NEW.updated_at = NOW();
RETURN NEW;
END;
$$;


SET default_tablespace = '';

SET default_table_access_method = heap;

--
-- TOC entry 223 (class 1259 OID 17168)
-- Name: Participant; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public."Participant" (
    id uuid NOT NULL,
    chat_id uuid NOT NULL,
    user_id integer NOT NULL,
    created_at timestamp without time zone DEFAULT now() NOT NULL,
    updated_at timestamp without time zone DEFAULT now() NOT NULL
);


--
-- TOC entry 219 (class 1259 OID 17075)
-- Name: alembic_version; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.alembic_version (
    version_num character varying(32) NOT NULL
);


--
-- TOC entry 220 (class 1259 OID 17130)
-- Name: chat; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.chat (
    id uuid NOT NULL,
    is_group boolean NOT NULL,
    last_message character varying,
    created_at timestamp without time zone DEFAULT now() NOT NULL,
    updated_at timestamp without time zone DEFAULT now() NOT NULL
);


--
-- TOC entry 224 (class 1259 OID 17195)
-- Name: message; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public.message (
    id uuid NOT NULL,
    sender_id integer NOT NULL,
    chat_id uuid NOT NULL,
    content character varying NOT NULL,
    is_read boolean NOT NULL,
    created_at timestamp without time zone DEFAULT now() NOT NULL,
    updated_at timestamp without time zone DEFAULT now() NOT NULL,
    is_edited boolean DEFAULT false NOT NULL
);


--
-- TOC entry 222 (class 1259 OID 17145)
-- Name: user; Type: TABLE; Schema: public; Owner: -
--

CREATE TABLE public."user" (
    id integer NOT NULL,
    display_name character varying NOT NULL,
    username character varying NOT NULL,
    email character varying NOT NULL,
    birth_day date,
    description character varying,
    avatar_url character varying,
    is_active boolean NOT NULL,
    is_verified boolean NOT NULL,
    is_superuser boolean NOT NULL,
    hashed_password character varying NOT NULL,
    created_at timestamp without time zone DEFAULT now() NOT NULL,
    updated_at timestamp without time zone DEFAULT now() NOT NULL
);


--
-- TOC entry 221 (class 1259 OID 17144)
-- Name: user_id_seq; Type: SEQUENCE; Schema: public; Owner: -
--

CREATE SEQUENCE public.user_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


--
-- TOC entry 3494 (class 0 OID 0)
-- Dependencies: 221
-- Name: user_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: -
--

ALTER SEQUENCE public.user_id_seq OWNED BY public."user".id;


--
-- TOC entry 3308 (class 2604 OID 17148)
-- Name: user id; Type: DEFAULT; Schema: public; Owner: -
--

ALTER TABLE ONLY public."user" ALTER COLUMN id SET DEFAULT nextval('public.user_id_seq'::regclass);


--
-- TOC entry 3329 (class 2606 OID 17179)
-- Name: Participant Participant_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public."Participant"
    ADD CONSTRAINT "Participant_pkey" PRIMARY KEY (id);


--
-- TOC entry 3317 (class 2606 OID 17080)
-- Name: alembic_version alembic_version_pkc; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.alembic_version
    ADD CONSTRAINT alembic_version_pkc PRIMARY KEY (version_num);


--
-- TOC entry 3320 (class 2606 OID 17142)
-- Name: chat chat_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.chat
    ADD CONSTRAINT chat_pkey PRIMARY KEY (id);


--
-- TOC entry 3336 (class 2606 OID 17210)
-- Name: message message_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.message
    ADD CONSTRAINT message_pkey PRIMARY KEY (id);


--
-- TOC entry 3332 (class 2606 OID 17181)
-- Name: Participant uq_chat_user; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public."Participant"
    ADD CONSTRAINT uq_chat_user UNIQUE (chat_id, user_id);


--
-- TOC entry 3324 (class 2606 OID 17164)
-- Name: user user_pkey; Type: CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public."user"
    ADD CONSTRAINT user_pkey PRIMARY KEY (id);


--
-- TOC entry 3326 (class 1259 OID 17192)
-- Name: Participant_chat_id_idx; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX "Participant_chat_id_idx" ON public."Participant" USING btree (chat_id);


--
-- TOC entry 3327 (class 1259 OID 17193)
-- Name: Participant_id_idx; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX "Participant_id_idx" ON public."Participant" USING btree (id);


--
-- TOC entry 3330 (class 1259 OID 17194)
-- Name: Participant_user_id_idx; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX "Participant_user_id_idx" ON public."Participant" USING btree (user_id);


--
-- TOC entry 3318 (class 1259 OID 17143)
-- Name: chat_id_idx; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX chat_id_idx ON public.chat USING btree (id);


--
-- TOC entry 3333 (class 1259 OID 17221)
-- Name: message_chat_id_idx; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX message_chat_id_idx ON public.message USING btree (chat_id);


--
-- TOC entry 3334 (class 1259 OID 17222)
-- Name: message_id_idx; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX message_id_idx ON public.message USING btree (id);


--
-- TOC entry 3337 (class 1259 OID 17223)
-- Name: message_sender_id_idx; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX message_sender_id_idx ON public.message USING btree (sender_id);


--
-- TOC entry 3321 (class 1259 OID 17165)
-- Name: user_email_idx; Type: INDEX; Schema: public; Owner: -
--

CREATE UNIQUE INDEX user_email_idx ON public."user" USING btree (email);


--
-- TOC entry 3322 (class 1259 OID 17166)
-- Name: user_id_idx; Type: INDEX; Schema: public; Owner: -
--

CREATE INDEX user_id_idx ON public."user" USING btree (id);


--
-- TOC entry 3325 (class 1259 OID 17167)
-- Name: user_username_idx; Type: INDEX; Schema: public; Owner: -
--

CREATE UNIQUE INDEX user_username_idx ON public."user" USING btree (username);


--
-- TOC entry 3338 (class 2606 OID 17182)
-- Name: Participant Participant_chat_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public."Participant"
    ADD CONSTRAINT "Participant_chat_id_fkey" FOREIGN KEY (chat_id) REFERENCES public.chat(id) ON DELETE CASCADE;


--
-- TOC entry 3339 (class 2606 OID 17187)
-- Name: Participant Participant_user_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public."Participant"
    ADD CONSTRAINT "Participant_user_id_fkey" FOREIGN KEY (user_id) REFERENCES public."user"(id) ON DELETE CASCADE;


--
-- TOC entry 3340 (class 2606 OID 17211)
-- Name: message message_chat_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.message
    ADD CONSTRAINT message_chat_id_fkey FOREIGN KEY (chat_id) REFERENCES public.chat(id) ON DELETE CASCADE;


--
-- TOC entry 3341 (class 2606 OID 17216)
-- Name: message message_sender_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: -
--

ALTER TABLE ONLY public.message
    ADD CONSTRAINT message_sender_id_fkey FOREIGN KEY (sender_id) REFERENCES public."user"(id) ON DELETE SET NULL;


