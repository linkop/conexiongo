PGDMP     1    5                {            goweb_db    15.2    15.2                0    0    ENCODING    ENCODING        SET client_encoding = 'UTF8';
                      false                       0    0 
   STDSTRINGS 
   STDSTRINGS     (   SET standard_conforming_strings = 'on';
                      false                       0    0 
   SEARCHPATH 
   SEARCHPATH     8   SELECT pg_catalog.set_config('search_path', '', false);
                      false                       1262    16410    goweb_db    DATABASE     ?   CREATE DATABASE goweb_db WITH TEMPLATE = template0 ENCODING = 'UTF8' LOCALE_PROVIDER = libc LOCALE = 'English_United States.1252';
    DROP DATABASE goweb_db;
                postgres    false            ?            1259    16435    libros    TABLE     ?   CREATE TABLE public.libros (
    codigo integer NOT NULL,
    titulo character varying(20),
    autor character varying(30),
    editorial character varying(15)
);
    DROP TABLE public.libros;
       public         heap    postgres    false            ?            1259    16434    libros_codigo_seq    SEQUENCE     ?   CREATE SEQUENCE public.libros_codigo_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;
 (   DROP SEQUENCE public.libros_codigo_seq;
       public          postgres    false    217                       0    0    libros_codigo_seq    SEQUENCE OWNED BY     G   ALTER SEQUENCE public.libros_codigo_seq OWNED BY public.libros.codigo;
          public          postgres    false    216            ?            1259    16427    users    TABLE     ?   CREATE TABLE public.users (
    id integer NOT NULL,
    username character varying(30) NOT NULL,
    password character varying(100) NOT NULL,
    email character varying(50)
);
    DROP TABLE public.users;
       public         heap    postgres    false            ?            1259    16426    users_id_seq    SEQUENCE     ?   CREATE SEQUENCE public.users_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;
 #   DROP SEQUENCE public.users_id_seq;
       public          postgres    false    215            	           0    0    users_id_seq    SEQUENCE OWNED BY     =   ALTER SEQUENCE public.users_id_seq OWNED BY public.users.id;
          public          postgres    false    214            k           2604    16438    libros codigo    DEFAULT     n   ALTER TABLE ONLY public.libros ALTER COLUMN codigo SET DEFAULT nextval('public.libros_codigo_seq'::regclass);
 <   ALTER TABLE public.libros ALTER COLUMN codigo DROP DEFAULT;
       public          postgres    false    217    216    217            j           2604    16430    users id    DEFAULT     d   ALTER TABLE ONLY public.users ALTER COLUMN id SET DEFAULT nextval('public.users_id_seq'::regclass);
 7   ALTER TABLE public.users ALTER COLUMN id DROP DEFAULT;
       public          postgres    false    214    215    215                      0    16435    libros 
   TABLE DATA           B   COPY public.libros (codigo, titulo, autor, editorial) FROM stdin;
    public          postgres    false    217   ?       ?          0    16427    users 
   TABLE DATA           >   COPY public.users (id, username, password, email) FROM stdin;
    public          postgres    false    215          
           0    0    libros_codigo_seq    SEQUENCE SET     @   SELECT pg_catalog.setval('public.libros_codigo_seq', 1, false);
          public          postgres    false    216                       0    0    users_id_seq    SEQUENCE SET     :   SELECT pg_catalog.setval('public.users_id_seq', 4, true);
          public          postgres    false    214            o           2606    16440    libros libros_pkey 
   CONSTRAINT     T   ALTER TABLE ONLY public.libros
    ADD CONSTRAINT libros_pkey PRIMARY KEY (codigo);
 <   ALTER TABLE ONLY public.libros DROP CONSTRAINT libros_pkey;
       public            postgres    false    217            m           2606    16432    users users_pkey 
   CONSTRAINT     N   ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_pkey PRIMARY KEY (id);
 :   ALTER TABLE ONLY public.users DROP CONSTRAINT users_pkey;
       public            postgres    false    215                  x?????? ? ?      ?   @   x?3?L?I? ?F?`?!=713G/9??ː3#?'?????B?2?????.?ȥ?f"I??qqq ?Q?     