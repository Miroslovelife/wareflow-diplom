CREATE TABLE public.permissions (
                                    id BIGSERIAL PRIMARY KEY,
                                    name VARCHAR(50) UNIQUE NOT NULL
);

CREATE TABLE public.roles (
                              id BIGSERIAL PRIMARY KEY,
                              name VARCHAR(50) NOT NULL
);

CREATE TABLE public.role_permissions (
                                         role_id INTEGER NOT NULL REFERENCES public.roles(id) ON DELETE CASCADE,
                                         permission_id INTEGER NOT NULL REFERENCES public.permissions(id) ON DELETE CASCADE,
                                         PRIMARY KEY (role_id, permission_id)
);

CREATE TABLE public.users (
                              uuid UUID DEFAULT gen_random_uuid() PRIMARY KEY,
                              phone_number VARCHAR(15),
                              username VARCHAR(50) UNIQUE,
                              first_name VARCHAR(50),
                              last_name VARCHAR(50),
                              surname VARCHAR(50),
                              email VARCHAR(100),
                              password VARCHAR(100),
                              role VARCHAR(20) DEFAULT 'user'
);

CREATE TABLE public.ware_houses (
                                    id BIGSERIAL PRIMARY KEY,
                                    address VARCHAR(200),
                                    name VARCHAR(100),
                                    uuid_user UUID NOT NULL REFERENCES public.users(uuid) ON DELETE CASCADE ON UPDATE CASCADE,
                                    capacity BIGINT,
                                    CONSTRAINT unique_user_id_name UNIQUE (uuid_user, name)
);

CREATE TABLE public.zones (
                              id BIGSERIAL PRIMARY KEY,
                              name VARCHAR(100) NOT NULL,
                              capacity BIGINT NOT NULL,
                              ware_house_id BIGINT NOT NULL REFERENCES public.ware_houses(id) ON DELETE CASCADE ON UPDATE CASCADE
);

CREATE TABLE public.products (
                                 uuid UUID DEFAULT gen_random_uuid() PRIMARY KEY,
                                 title CHAR(200),
                                 count BIGINT NOT NULL,
                                 qr CHAR(500),
                                 description CHAR(500),
                                 zone_id BIGINT NOT NULL REFERENCES public.zones(id) ON DELETE CASCADE ON UPDATE CASCADE
);

CREATE TABLE public.warehouse_user_roles (
                                             ware_house_id INTEGER NOT NULL REFERENCES public.ware_houses(id) ON DELETE CASCADE,
                                             user_id UUID NOT NULL REFERENCES public.users(uuid) ON DELETE CASCADE,
                                             role_id INTEGER NOT NULL REFERENCES public.roles(id) ON DELETE CASCADE,
                                             PRIMARY KEY (ware_house_id, user_id, role_id)
);
