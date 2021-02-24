--
-- PostgreSQL database dump
--

-- Dumped from database version 13.1
-- Dumped by pg_dump version 13.1

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
-- Name: CallBack; Type: TABLE; Schema: public; Owner: hivemind
--

CREATE TABLE public."CallBack" (
    "UUIDImplant" text NOT NULL,
    "FirstCall" text NOT NULL,
    "LastCall" text
);


ALTER TABLE public."CallBack" OWNER TO hivemind;

--
-- Name: ExecutedActions; Type: TABLE; Schema: public; Owner: hivemind
--

CREATE TABLE public."ExecutedActions" (
    "id" text NOT NULL,
    "UUIDofAction" text NOT NULL,
    "TimeSent" text NOT NULL,
    "TimeRan" text,
    "Successful" boolean,
    "ActionResponse" text
);


ALTER TABLE public."ExecutedActions" OWNER TO hivemind;

--
-- Name: Groups; Type: TABLE; Schema: public; Owner: hivemind
--

CREATE TABLE public."Groups" (
    "UUID" text NOT NULL,
    "GroupName" text NOT NULL,
    "Implants" text[]
);


ALTER TABLE public."Groups" OWNER TO hivemind;

--
-- Name: Implant; Type: TABLE; Schema: public; Owner: hivemind
--

CREATE TABLE public."Implant" (
    "UUID" text NOT NULL,
    "UUIDImplantType" text NOT NULL,
    "PrimaryIP" text NOT NULL,
    "Hostname" text,
    "MAC" text,
    "ImplantOS" text,
    "OtherIPs" text[],
    "SupportedModules" text[]
);


ALTER TABLE public."Implant" OWNER TO hivemind;

--
-- Name: ImplantType; Type: TABLE; Schema: public; Owner: hivemind
--

CREATE TABLE public."ImplantType" (
    "UUID" text NOT NULL,
    "ImplantName" text NOT NULL,
    "ImplantVersion" text NOT NULL
);


ALTER TABLE public."ImplantType" OWNER TO hivemind;

--
-- Name: ModuleFuncs; Type: TABLE; Schema: public; Owner: hivemind
--

CREATE TABLE public."ModuleFuncs" (
    "UUID" text NOT NULL,
    "ModuleFuncName" text NOT NULL,
    "ModuleFuncDesc" text NOT NULL,
    "NumOfParameters" integer NOT NULL,
    "ParameterTypes" text[] NOT NULL,
    "ParameterNames" text[] NOT NULL
);


ALTER TABLE public."ModuleFuncs" OWNER TO hivemind;

--
-- Name: Modules; Type: TABLE; Schema: public; Owner: hivemind
--

CREATE TABLE public."Modules" (
    "ModuleName" text NOT NULL,
    "ModuleDesc" text NOT NULL,
    "ModuleFuncIds" text[] NOT NULL
);


ALTER TABLE public."Modules" OWNER TO hivemind;

--
-- Name: Operators; Type: TABLE; Schema: public; Owner: hivemind
--

CREATE TABLE public."Operators" (
    "Username" text NOT NULL,
    "Password" text NOT NULL,
    "Permission" integer NOT NULL
);


ALTER TABLE public."Operators" OWNER TO hivemind;

--
-- Name: ParamTypes; Type: TABLE; Schema: public; Owner: hivemind
--

CREATE TABLE public."ParamTypes" (
    "TypeName" text NOT NULL,
    "IsComboOption" boolean,
    "ComboOptions" text[]
);


ALTER TABLE public."ParamTypes" OWNER TO hivemind;

--
-- Name: StagedActions; Type: TABLE; Schema: public; Owner: hivemind
--

CREATE TABLE public."StagedActions" (
    id text NOT NULL,
    "UUIDofAction" text NOT NULL,
    "UUIDofImplant" text NOT NULL,
    "TimeStaged" text NOT NULL
);


ALTER TABLE public."StagedActions" OWNER TO hivemind;

--
-- Name: StoredActions; Type: TABLE; Schema: public; Owner: hivemind
--

CREATE TABLE public."StoredActions" (
    "UUID" text NOT NULL,
    "ModuleToRun" text NOT NULL,
    "ModuleFunc" text NOT NULL,
    "Arguments" text[]
);


ALTER TABLE public."StoredActions" OWNER TO hivemind;

--
-- Data for Name: CallBack; Type: TABLE DATA; Schema: public; Owner: hivemind
--

COPY public."CallBack" ("UUIDImplant", "FirstCall", "LastCall") FROM stdin;
\.


--
-- Data for Name: ExecutedActions; Type: TABLE DATA; Schema: public; Owner: hivemind
--

COPY public."ExecutedActions" (id, "UUIDofAction", "TimeSent", "TimeRan", "Successful", "ActionResponse") FROM stdin;
\.


--
-- Data for Name: Groups; Type: TABLE DATA; Schema: public; Owner: hivemind
--

COPY public."Groups" ("UUID", "GroupName", "Implants") FROM stdin;
\.


--
-- Data for Name: Implant; Type: TABLE DATA; Schema: public; Owner: hivemind
--

COPY public."Implant" ("UUID", "UUIDImplantType", "PrimaryIP", "Hostname", "MAC", "ImplantOS", "OtherIPs", "SupportedModules") FROM stdin;
\.


--
-- Data for Name: ImplantType; Type: TABLE DATA; Schema: public; Owner: hivemind
--

COPY public."ImplantType" ("UUID", "ImplantName", "ImplantVersion") FROM stdin;
\.


--
-- Data for Name: ModuleFuncs; Type: TABLE DATA; Schema: public; Owner: hivemind
--

COPY public."ModuleFuncs" ("UUID", "ModuleFuncName", "NumOfParameters", "ParameterTypes", "ParameterNames") FROM stdin;
\.


--
-- Data for Name: Modules; Type: TABLE DATA; Schema: public; Owner: hivemind
--

COPY public."Modules" ("ModuleName", "ModuleFuncIds") FROM stdin;
\.


--
-- Data for Name: Operators; Type: TABLE DATA; Schema: public; Owner: hivemind
--

COPY public."Operators" ("Username", "Password", "Permission") FROM stdin;
\.


--
-- Data for Name: ParamTypes; Type: TABLE DATA; Schema: public; Owner: hivemind
--

COPY public."ParamTypes" ("TypeName", "IsComboOption", "ComboOptions") FROM stdin;
\.


--
-- Data for Name: StagedActions; Type: TABLE DATA; Schema: public; Owner: hivemind
--

COPY public."StagedActions" (id, "UUIDofAction", "UUIDofImplant", "TimeStaged") FROM stdin;
\.


--
-- Data for Name: StoredActions; Type: TABLE DATA; Schema: public; Owner: hivemind
--

COPY public."StoredActions" ("UUID", "ModuleToRun", "ModuleFunc", "Arguments") FROM stdin;
\.


--
-- Name: CallBack CallBack_pkey; Type: CONSTRAINT; Schema: public; Owner: hivemind
--

ALTER TABLE ONLY public."CallBack"
    ADD CONSTRAINT "CallBack_pkey" PRIMARY KEY ("UUIDImplant");


--
-- Name: ExecutedActions ExecutedActions_pkey; Type: CONSTRAINT; Schema: public; Owner: hivemind
--

ALTER TABLE ONLY public."ExecutedActions"
    ADD CONSTRAINT "ExecutedActions_pkey" PRIMARY KEY (id);


--
-- Name: Groups Groups_pkey; Type: CONSTRAINT; Schema: public; Owner: hivemind
--

ALTER TABLE ONLY public."Groups"
    ADD CONSTRAINT "Groups_pkey" PRIMARY KEY ("UUID");


--
-- Name: ImplantType ImplantType_pkey; Type: CONSTRAINT; Schema: public; Owner: hivemind
--

ALTER TABLE ONLY public."ImplantType"
    ADD CONSTRAINT "ImplantType_pkey" PRIMARY KEY ("UUID");


--
-- Name: Implant Implant_pkey; Type: CONSTRAINT; Schema: public; Owner: hivemind
--

ALTER TABLE ONLY public."Implant"
    ADD CONSTRAINT "Implant_pkey" PRIMARY KEY ("UUID");


--
-- Name: ModuleFuncs ModuleFuncs_pkey; Type: CONSTRAINT; Schema: public; Owner: hivemind
--

ALTER TABLE ONLY public."ModuleFuncs"
    ADD CONSTRAINT "ModuleFuncs_pkey" PRIMARY KEY ("UUID");


--
-- Name: Modules Modules_pkey; Type: CONSTRAINT; Schema: public; Owner: hivemind
--

ALTER TABLE ONLY public."Modules"
    ADD CONSTRAINT "Modules_pkey" PRIMARY KEY ("ModuleName");


--
-- Name: Operators Operators_pkey; Type: CONSTRAINT; Schema: public; Owner: hivemind
--

ALTER TABLE ONLY public."Operators"
    ADD CONSTRAINT "Operators_pkey" PRIMARY KEY ("Username");


--
-- Name: ParamTypes ParamTypes_pkey; Type: CONSTRAINT; Schema: public; Owner: hivemind
--

ALTER TABLE ONLY public."ParamTypes"
    ADD CONSTRAINT "ParamTypes_pkey" PRIMARY KEY ("TypeName");


--
-- Name: StagedActions StagedActions_pkey; Type: CONSTRAINT; Schema: public; Owner: hivemind
--

ALTER TABLE ONLY public."StagedActions"
    ADD CONSTRAINT "StagedActions_pkey" PRIMARY KEY (id);


--
-- Name: StoredActions StoredActions_pkey; Type: CONSTRAINT; Schema: public; Owner: hivemind
--

ALTER TABLE ONLY public."StoredActions"
    ADD CONSTRAINT "StoredActions_pkey" PRIMARY KEY ("UUID");


--
-- Name: CallBack UUIDImplant; Type: FK CONSTRAINT; Schema: public; Owner: hivemind
--

ALTER TABLE ONLY public."CallBack"
    ADD CONSTRAINT "UUIDImplant" FOREIGN KEY ("UUIDImplant") REFERENCES public."Implant"("UUID") NOT VALID;


--
-- Name: StagedActions implant_fk; Type: FK CONSTRAINT; Schema: public; Owner: hivemind
--

ALTER TABLE ONLY public."StagedActions"
    ADD CONSTRAINT implant_fk FOREIGN KEY ("UUIDofImplant") REFERENCES public."Implant"("UUID");


--
-- Name: StagedActions storedaction_fk; Type: FK CONSTRAINT; Schema: public; Owner: hivemind
--

ALTER TABLE ONLY public."StagedActions"
    ADD CONSTRAINT storedaction_fk FOREIGN KEY ("UUIDofAction") REFERENCES public."StoredActions"("UUID");


--
-- Name: ExecutedActions storedactions_fk; Type: FK CONSTRAINT; Schema: public; Owner: hivemind
--

ALTER TABLE ONLY public."ExecutedActions"
    ADD CONSTRAINT storedactions_fk FOREIGN KEY ("UUIDofAction") REFERENCES public."StoredActions"("UUID") NOT VALID;


--
-- Name: Implant uuid_fk; Type: FK CONSTRAINT; Schema: public; Owner: hivemind
--

ALTER TABLE ONLY public."Implant"
    ADD CONSTRAINT uuid_fk FOREIGN KEY ("UUIDImplantType") REFERENCES public."ImplantType"("UUID") NOT VALID;


--
-- PostgreSQL database dump complete
--

