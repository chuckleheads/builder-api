CREATE TYPE origin_package_operation AS ENUM (
    'promote',
    'demote'
);

CREATE TYPE origin_package_visibility AS ENUM (
    'public',
    'private',
    'hidden'
);

CREATE TYPE package_channel_trigger AS ENUM (
    'unknown',
    'builder_ui',
    'hab_client'
);

CREATE SEQUENCE groups_id_seq;
CREATE SEQUENCE group_projects_id_seq;
CREATE SEQUENCE origin_package_id_seq;
CREATE SEQUENCE job_id_seq;
CREATE SEQUENCE account_tokens_id_seq;
CREATE SEQUENCE accounts_id_seq;
CREATE SEQUENCE origin_channel_id_seq;
CREATE SEQUENCE origin_id_seq;
CREATE SEQUENCE origin_integration_id_seq;
CREATE SEQUENCE origin_invitations_id_seq;
CREATE SEQUENCE origin_private_encryption_key_id_seq;
CREATE SEQUENCE origin_project_id_seq;
CREATE SEQUENCE origin_project_integration_id_seq;
CREATE SEQUENCE origin_public_encryption_key_id_seq;
CREATE SEQUENCE origin_public_key_id_seq;
CREATE SEQUENCE origin_secret_key_id_seq;
CREATE SEQUENCE origin_secrets_id_seq;


CREATE FUNCTION next_id_v1(sequence_id regclass, OUT result bigint) RETURNS bigint
    LANGUAGE plpgsql
    AS $$
                DECLARE
                    our_epoch bigint := 1409266191000;
                    seq_id bigint;
                    now_millis bigint;
                BEGIN
                    SELECT nextval(sequence_id) % 1024 INTO seq_id;
                    SELECT FLOOR(EXTRACT(EPOCH FROM clock_timestamp()) * 1000) INTO now_millis;
                    result := (now_millis - our_epoch) << 23;
                    result := result | (seq_id << 13);
                END;
                $$;

CREATE TABLE origin_packages (
    id bigint DEFAULT next_id_v1('origin_package_id_seq'::regclass) NOT NULL,
    owner_id bigint,
    name text,
    ident text,
    ident_array text[],
    checksum text,
    manifest text,
    config text,
    target text,
    deps text[],
    tdeps text[],
    exposes integer[],
    scheduler_sync boolean DEFAULT false,
    created_at timestamp with time zone DEFAULT now(),
    updated_at timestamp with time zone DEFAULT now(),
    visibility origin_package_visibility DEFAULT 'public'::origin_package_visibility NOT NULL,
    ident_vector tsvector,
    origin text
);

CREATE TABLE groups (
    id bigint DEFAULT next_id_v1('groups_id_seq'::regclass) NOT NULL,
    group_state text,
    project_name text,
    created_at timestamp with time zone DEFAULT now(),
    updated_at timestamp with time zone DEFAULT now(),
    target text DEFAULT 'x86_64-linux'::text
);


CREATE TABLE group_projects (
    id bigint DEFAULT nextval('group_projects_id_seq'::regclass) NOT NULL,
    owner_id bigint,
    project_name text,
    project_ident text,
    project_state text,
    job_id bigint DEFAULT 0,
    created_at timestamp with time zone DEFAULT now(),
    updated_at timestamp with time zone DEFAULT now(),
    target text DEFAULT 'x86_64-linux'::text
);

CREATE TABLE busy_workers (
    id serial NOT NULL PRIMARY KEY,
    ident text,
    job_id bigint,
    quarantined boolean,
    created_at timestamp with time zone DEFAULT now(),
    updated_at timestamp with time zone DEFAULT now(),
    target text DEFAULT 'x86_64-linux'::text
);

CREATE TABLE jobs (
    id bigint DEFAULT next_id_v1('job_id_seq'::regclass) NOT NULL,
    owner_id bigint,
    job_state text DEFAULT 'Pending'::text,
    project_id bigint,
    project_name text,
    project_owner_id bigint,
    project_plan_path text,
    vcs text,
    vcs_arguments text[],
    net_error_code integer,
    net_error_msg text,
    scheduler_sync boolean DEFAULT false,
    created_at timestamp with time zone DEFAULT now(),
    updated_at timestamp with time zone DEFAULT now(),
    build_started_at timestamp with time zone,
    build_finished_at timestamp with time zone,
    package_ident text,
    archived boolean DEFAULT false NOT NULL,
    channel text,
    sync_count integer DEFAULT 0,
    worker text,
    target text DEFAULT 'x86_64-linux'::text
);

CREATE TABLE account_tokens (
    id bigint DEFAULT next_id_v1('account_tokens_id_seq'::regclass) NOT NULL,
    account_id bigint,
    token text,
    created_at timestamp with time zone DEFAULT now()
);

CREATE TABLE accounts (
    id bigint DEFAULT next_id_v1('accounts_id_seq'::regclass) NOT NULL,
    name text,
    email text,
    created_at timestamp with time zone DEFAULT now(),
    updated_at timestamp with time zone DEFAULT now()
);

CREATE TABLE audit_package (
    id serial NOT NULL PRIMARY KEY,
    operation origin_package_operation,
    trigger package_channel_trigger,
    requester_id bigint,
    requester_name text,
    created_at timestamp with time zone DEFAULT now(),
    origin text,
    channel text,
    package_ident text
);

CREATE TABLE audit_package_group (
    id SERIAL NOT NULL PRIMARY KEY,
    package_ids bigint[],
    operation origin_package_operation,
    trigger package_channel_trigger,
    requester_id bigint,
    requester_name text,
    group_id bigint,
    created_at timestamp with time zone DEFAULT now(),
    origin text,
    channel text
);

CREATE TABLE origin_channel_packages (
    channel_id bigint NOT NULL,
    package_id bigint NOT NULL,
    created_at timestamp with time zone DEFAULT now(),
    updated_at timestamp with time zone DEFAULT now()
);

CREATE TABLE origin_channels (
    id bigint DEFAULT next_id_v1('origin_channel_id_seq'::regclass) NOT NULL,
    owner_id bigint,
    name text,
    created_at timestamp with time zone DEFAULT now(),
    updated_at timestamp with time zone DEFAULT now(),
    origin text
);



CREATE TABLE origin_integrations (
    id bigint DEFAULT next_id_v1('origin_integration_id_seq'::regclass) NOT NULL,
    origin text,
    integration text,
    name text,
    body text,
    created_at timestamp with time zone DEFAULT now(),
    updated_at timestamp with time zone DEFAULT now()
);



CREATE TABLE origin_invitations (
    id bigint DEFAULT next_id_v1('origin_invitations_id_seq'::regclass) NOT NULL,
    origin text,
    account_id bigint,
    account_name text,
    owner_id bigint,
    ignored boolean DEFAULT false,
    created_at timestamp with time zone DEFAULT now(),
    updated_at timestamp with time zone DEFAULT now()
);

CREATE TABLE origin_members (
    id serial NOT NULL PRIMARY KEY,
    origin text,
    account_id bigint NOT NULL,
    created_at timestamp with time zone DEFAULT now(),
    updated_at timestamp with time zone DEFAULT now()
);

CREATE VIEW origin_package_versions AS
 SELECT origin_packages.origin,
    origin_packages.name,
    origin_packages.visibility,
    origin_packages.ident_array[3] AS version,
    count(origin_packages.ident_array[4]) AS release_count,
    max(origin_packages.ident_array[4]) AS latest,
    array_agg(DISTINCT origin_packages.target) AS platforms
   FROM origin_packages
  GROUP BY origin_packages.ident_array[3], origin_packages.origin, origin_packages.name, origin_packages.visibility;

CREATE TABLE origin_private_encryption_keys (
    id bigint DEFAULT next_id_v1('origin_private_encryption_key_id_seq'::regclass) NOT NULL,
    owner_id bigint,
    name text,
    revision text,
    full_name text,
    body bytea,
    created_at timestamp with time zone DEFAULT now(),
    updated_at timestamp with time zone DEFAULT now(),
    origin text
);


CREATE TABLE origin_project_integrations (
    id bigint DEFAULT next_id_v1('origin_project_integration_id_seq'::regclass) NOT NULL,
    origin text NOT NULL,
    body text NOT NULL,
    created_at timestamp with time zone DEFAULT now() NOT NULL,
    updated_at timestamp with time zone DEFAULT now() NOT NULL,
    project_id bigint NOT NULL,
    integration_id bigint NOT NULL
);

CREATE TABLE origin_projects (
    id bigint DEFAULT next_id_v1('origin_project_id_seq'::regclass) NOT NULL,
    origin text,
    package_name text,
    name text,
    plan_path text,
    owner_id bigint,
    vcs_type text,
    vcs_data text,
    created_at timestamp with time zone DEFAULT now(),
    updated_at timestamp with time zone DEFAULT now(),
    vcs_installation_id bigint,
    visibility origin_package_visibility DEFAULT 'public'::origin_package_visibility NOT NULL,
    auto_build boolean DEFAULT true NOT NULL
);


CREATE TABLE origin_public_encryption_keys (
    id bigint DEFAULT next_id_v1('origin_public_key_id_seq'::regclass) NOT NULL,
    owner_id bigint,
    name text,
    revision text,
    full_name text,
    body bytea,
    created_at timestamp with time zone DEFAULT now(),
    updated_at timestamp with time zone DEFAULT now(),
    origin text
);

CREATE TABLE origin_public_keys (
    id bigint DEFAULT next_id_v1('origin_public_key_id_seq'::regclass) NOT NULL,
    owner_id bigint,
    name text,
    revision text,
    full_name text,
    body bytea,
    created_at timestamp with time zone DEFAULT now(),
    updated_at timestamp with time zone DEFAULT now(),
    origin text
);


CREATE TABLE origin_secret_keys (
    id bigint DEFAULT next_id_v1('origin_secret_key_id_seq'::regclass) NOT NULL,
    owner_id bigint,
    name text,
    revision text,
    full_name text,
    body bytea,
    created_at timestamp with time zone DEFAULT now(),
    updated_at timestamp with time zone DEFAULT now(),
    origin text
);

CREATE TABLE origin_secrets (
    id bigint DEFAULT next_id_v1('origin_secrets_id_seq'::regclass) NOT NULL,
    owner_id bigint,
    name text,
    value text,
    created_at timestamp with time zone DEFAULT now(),
    updated_at timestamp with time zone DEFAULT now(),
    origin text
);

CREATE TABLE origins (
    name text NOT NULL,
    owner_id bigint,
    created_at timestamp with time zone DEFAULT now(),
    updated_at timestamp with time zone DEFAULT now(),
    default_package_visibility origin_package_visibility DEFAULT 'public'::origin_package_visibility NOT NULL
);

CREATE VIEW origins_with_secret_key AS
 SELECT origins.name,
    origins.owner_id,
    origin_secret_keys.full_name AS private_key_name,
    origins.default_package_visibility
   FROM (origins
     LEFT JOIN origin_secret_keys ON ((origins.name = origin_secret_keys.origin)))
  ORDER BY origins.name, origin_secret_keys.full_name DESC;

CREATE TABLE origins_with_stats (
    name text,
    owner_id bigint,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    default_package_visibility origin_package_visibility,
    package_count bigint
);

CREATE TABLE audit_jobs (
    id serial NOT NULL PRIMARY KEY,
    group_id bigint,
    operation smallint,
    trigger smallint,
    requester_id bigint,
    requester_name text,
    created_at timestamp with time zone DEFAULT now()
);

CREATE FUNCTION abort_group_v1(in_gid bigint) RETURNS void
    LANGUAGE sql
    AS $$
  UPDATE group_projects SET project_state='Failure'
    WHERE owner_id = in_gid
    AND (project_state = 'InProgress' OR project_state = 'NotStarted');
  UPDATE groups SET group_state='Complete' where id = in_gid;
$$;


CREATE FUNCTION add_audit_jobs_entry_v1(p_group_id bigint, p_operation smallint, p_trigger smallint, p_requester_id bigint, p_requester_name text) RETURNS SETOF audit_jobs
    LANGUAGE sql
    AS $$
      INSERT INTO audit_jobs (group_id, operation, trigger, requester_id, requester_name)
      VALUES (p_group_id, p_operation, p_trigger, p_requester_id, p_requester_name)
      RETURNING *;
$$;

CREATE FUNCTION cancel_group_v1(in_gid bigint) RETURNS void
    LANGUAGE sql
    AS $$
  UPDATE group_projects SET project_state='Canceled'
    WHERE owner_id = in_gid
    AND (project_state = 'NotStarted');
  UPDATE groups SET group_state='Canceled' where id = in_gid;
$$;

CREATE FUNCTION check_active_group_v1(pname text) RETURNS SETOF groups
    LANGUAGE sql
    AS $$
  SELECT * FROM groups
  WHERE project_name = pname
  AND group_state IN ('Pending', 'Dispatching')
$$;

CREATE FUNCTION count_group_projects_v2(origin text) RETURNS bigint
    LANGUAGE plpgsql STABLE
    AS $$
  BEGIN
    RETURN COUNT(*) FROM group_projects WHERE project_ident LIKE (origin || '/%');
  END
$$;

CREATE FUNCTION count_jobs_v1(in_job_state text) RETURNS bigint
    LANGUAGE plpgsql STABLE
    AS $$
  BEGIN
    RETURN COUNT(*) FROM jobs WHERE job_state = in_job_state;
  END
$$;

CREATE FUNCTION count_origin_packages_v1(origin text) RETURNS bigint
    LANGUAGE plpgsql STABLE
    AS $$
BEGIN
  RETURN COUNT(*) FROM origin_packages WHERE ident_array[1]=origin;
END
$$;

CREATE FUNCTION count_unique_origin_packages_v1(op_origin text) RETURNS bigint
    LANGUAGE sql STABLE
    AS $$
  SELECT COUNT(DISTINCT ident_array[2]) AS total
  FROM origin_packages
  WHERE ident_array[1] = op_origin
$$;

CREATE FUNCTION delete_busy_worker_v1(in_ident text, in_job_id bigint) RETURNS void
    LANGUAGE sql
    AS $$
  DELETE FROM busy_workers
  WHERE ident = in_ident AND job_id = in_job_id
$$;

CREATE FUNCTION find_group_project_v1(gid bigint, name text) RETURNS SETOF group_projects
    LANGUAGE plpgsql STABLE
    AS $$
BEGIN
  RETURN QUERY SELECT * FROM group_projects WHERE owner_id = gid AND project_name = name;
  RETURN;
END
$$;

CREATE FUNCTION get_all_origin_packages_for_ident_v1(op_ident text) RETURNS SETOF origin_packages
    LANGUAGE plpgsql STABLE
    AS $$
  BEGIN
    RETURN QUERY SELECT * FROM origin_packages WHERE ident LIKE (op_ident || '%') ORDER BY ident;
    RETURN;
  END
  $$;

CREATE FUNCTION get_busy_workers_v1() RETURNS SETOF busy_workers
    LANGUAGE sql STABLE
    AS $$
  SELECT * FROM busy_workers
$$;

CREATE FUNCTION get_cancel_pending_jobs_v1() RETURNS SETOF jobs
    LANGUAGE sql
    AS $$
  SELECT *
  FROM jobs
  WHERE job_state = 'CancelPending'
$$;

CREATE FUNCTION get_dispatched_jobs_v1() RETURNS SETOF jobs
    LANGUAGE sql STABLE
    AS $$
  SELECT *
  FROM jobs
  WHERE job_state = 'Dispatched'
$$;

CREATE FUNCTION get_group_projects_for_group_v1(gid bigint) RETURNS SETOF group_projects
    LANGUAGE plpgsql STABLE
    AS $$
  BEGIN
    RETURN QUERY SELECT * FROM group_projects WHERE owner_id = gid;
    RETURN;
  END
$$;

CREATE FUNCTION get_group_v1(gid bigint) RETURNS SETOF groups
    LANGUAGE plpgsql STABLE
    AS $$
BEGIN
  RETURN QUERY SELECT * FROM groups WHERE id = gid;
  RETURN;
END
$$;

CREATE FUNCTION get_job_groups_for_origin_v2(op_origin text, op_limit integer) RETURNS SETOF groups
    LANGUAGE sql STABLE
    AS $$
  SELECT *
  FROM groups
  WHERE project_name LIKE (op_origin || '/%')
  ORDER BY created_at DESC
  LIMIT op_limit
$$;

CREATE FUNCTION get_job_v1(jid bigint) RETURNS SETOF jobs
    LANGUAGE plpgsql STABLE
    AS $$
BEGIN
  RETURN QUERY SELECT * FROM jobs WHERE id = jid;
  RETURN;
END
$$;

CREATE FUNCTION get_jobs_for_project_v2(p_project_name text, p_limit bigint, p_offset bigint) RETURNS TABLE(total_count bigint, id bigint, owner_id bigint, job_state text, created_at timestamp with time zone, build_started_at timestamp with time zone, build_finished_at timestamp with time zone, package_ident text, project_id bigint, project_name text, project_owner_id bigint, project_plan_path text, vcs text, vcs_arguments text[], net_error_msg text, net_error_code integer, archived boolean)
    LANGUAGE sql STABLE
    AS $$
  SELECT COUNT(*) OVER () AS total_count, id, owner_id, job_state, created_at, build_started_at,
  build_finished_at, package_ident, project_id, project_name, project_owner_id, project_plan_path, vcs,
  vcs_arguments, net_error_msg, net_error_code, archived
  FROM jobs
  WHERE project_name = p_project_name
  ORDER BY created_at DESC
  LIMIT p_limit
  OFFSET p_offset;
$$;

CREATE FUNCTION get_origin_package_v5(op_ident text, op_visibilities origin_package_visibility[]) RETURNS SETOF origin_packages
    LANGUAGE sql STABLE
    AS $$
    SELECT *
    FROM origin_packages
    WHERE ident = op_ident
    AND visibility = ANY(op_visibilities);
$$;

CREATE FUNCTION get_queued_group_v1(pname text) RETURNS SETOF groups
    LANGUAGE sql
    AS $$
  SELECT * FROM groups
  WHERE project_name = pname
  AND group_state = 'Queued'
$$;

CREATE FUNCTION get_queued_groups_v1() RETURNS SETOF groups
    LANGUAGE sql
    AS $$
  SELECT * FROM groups
  WHERE group_state = 'Queued'
$$;

CREATE FUNCTION insert_group_v2(root_project text, project_names text[], project_idents text[]) RETURNS SETOF groups
    LANGUAGE sql
    AS $$
  WITH my_group AS (
          INSERT INTO groups (project_name, group_state)
          VALUES (root_project, 'Queued') RETURNING *
      ), my_project AS (
          INSERT INTO group_projects (owner_id, project_name, project_ident, project_state)
          SELECT g.id, project_info.name, project_info.ident, 'NotStarted'
          FROM my_group AS g, unnest(project_names, project_idents) AS project_info(name, ident)
      )
  SELECT * FROM my_group;
$$;

CREATE FUNCTION insert_group_v3(root_project text, project_names text[], project_idents text[], p_target text) RETURNS SETOF groups
    LANGUAGE sql
    AS $$
  WITH my_group AS (
          INSERT INTO groups (project_name, group_state, target)
          VALUES (root_project, 'Queued', p_target) RETURNING *
      ), my_project AS (
          INSERT INTO group_projects (owner_id, project_name, project_ident, project_state)
          SELECT g.id, project_info.name, project_info.ident, 'NotStarted'
          FROM my_group AS g, unnest(project_names, project_idents) AS project_info(name, ident)
      )
  SELECT * FROM my_group;
$$;

CREATE FUNCTION insert_job_v2(p_owner_id bigint, p_project_id bigint, p_project_name text, p_project_owner_id bigint, p_project_plan_path text, p_vcs text, p_vcs_arguments text[], p_channel text) RETURNS SETOF jobs
    LANGUAGE sql
    AS $$
      INSERT INTO jobs (owner_id, job_state, project_id, project_name, project_owner_id, project_plan_path, vcs, vcs_arguments, channel)
      VALUES (p_owner_id, 'Pending', p_project_id, p_project_name, p_project_owner_id, p_project_plan_path, p_vcs, p_vcs_arguments, p_channel)
      RETURNING *;
$$;

CREATE FUNCTION insert_job_v3(p_owner_id bigint, p_project_id bigint, p_project_name text, p_project_owner_id bigint, p_project_plan_path text, p_vcs text, p_vcs_arguments text[], p_channel text, p_target text) RETURNS SETOF jobs
    LANGUAGE sql
    AS $$
      INSERT INTO jobs (owner_id, job_state, project_id, project_name, project_owner_id, project_plan_path, vcs, vcs_arguments, channel, target)
      VALUES (p_owner_id, 'Pending', p_project_id, p_project_name, p_project_owner_id, p_project_plan_path, p_vcs, p_vcs_arguments, p_channel, p_target)
      RETURNING *;
$$;

CREATE FUNCTION mark_as_archived_v1(p_job_id bigint) RETURNS void
    LANGUAGE sql
    AS $$
  UPDATE jobs
  SET archived = TRUE
  WHERE id = p_job_id;
$$;

CREATE FUNCTION next_pending_job_v1(p_worker text) RETURNS SETOF jobs
    LANGUAGE plpgsql
    AS $$
DECLARE
    r jobs % rowtype;
BEGIN
    FOR r IN
        SELECT * FROM jobs
        WHERE job_state = 'Pending'
        ORDER BY created_at ASC
        FOR UPDATE SKIP LOCKED
        LIMIT 1
    LOOP
        UPDATE jobs SET job_state='Dispatched', scheduler_sync=false, worker=p_worker, updated_at=now()
        WHERE id=r.id
        RETURNING * INTO r;
        RETURN NEXT r;
    END LOOP;
  RETURN;
END
$$;

CREATE FUNCTION next_pending_job_v2(p_worker text, p_target text) RETURNS SETOF jobs
    LANGUAGE plpgsql
    AS $$
DECLARE
    r jobs % rowtype;
BEGIN
    FOR r IN
        SELECT * FROM jobs
        WHERE job_state = 'Pending' AND target = p_target
        ORDER BY created_at ASC
        FOR UPDATE SKIP LOCKED
        LIMIT 1
    LOOP
        UPDATE jobs SET job_state='Dispatched', scheduler_sync=false, worker=p_worker, updated_at=now()
        WHERE id=r.id
        RETURNING * INTO r;
        RETURN NEXT r;
    END LOOP;
  RETURN;
END
$$;

CREATE FUNCTION pending_groups_v1(integer) RETURNS SETOF groups
    LANGUAGE plpgsql
    AS $_$
DECLARE
    r groups % rowtype;
BEGIN
    FOR r IN
        SELECT * FROM groups
        WHERE group_state = 'Pending'
        ORDER BY created_at ASC
        FOR UPDATE SKIP LOCKED
        LIMIT $1
    LOOP
        UPDATE groups SET group_state='Dispatching', updated_at=now() WHERE id=r.id RETURNING * INTO r;
        RETURN NEXT r;
    END LOOP;
  RETURN;
END
$_$;

CREATE FUNCTION pending_jobs_v1(integer) RETURNS SETOF jobs
    LANGUAGE plpgsql
    AS $_$
DECLARE
  r jobs % rowtype;
BEGIN
  FOR r IN
    SELECT * FROM jobs
    WHERE job_state = 'Pending'
    ORDER BY created_at ASC
    FOR UPDATE SKIP LOCKED
    LIMIT $1
  LOOP
    UPDATE jobs SET job_state='Dispatched', scheduler_sync=false, updated_at=now() WHERE id=r.id RETURNING * INTO r;
    RETURN NEXT r;
  END LOOP;
  RETURN;
END
$_$;

CREATE FUNCTION set_group_project_name_state_v1(gid bigint, pname text, state text) RETURNS void
    LANGUAGE plpgsql
    AS $$
  BEGIN
    UPDATE group_projects SET project_state=state, updated_at=now() WHERE owner_id=gid AND project_name=pname;
  END
$$;

CREATE FUNCTION set_group_project_state_ident_v1(pid bigint, jid bigint, state text, ident text) RETURNS void
    LANGUAGE sql
    AS $$
  UPDATE group_projects SET project_state=state, job_id=jid, project_ident=ident, updated_at=now() WHERE id=pid;
$$;

CREATE FUNCTION set_group_project_state_v1(pid bigint, jid bigint, state text) RETURNS void
    LANGUAGE plpgsql
    AS $$
  BEGIN
    UPDATE group_projects SET project_state=state, job_id=jid, updated_at=now() WHERE id=pid;
  END
$$;

CREATE FUNCTION set_group_state_v1(gid bigint, gstate text) RETURNS void
    LANGUAGE plpgsql
    AS $$
  BEGIN
      UPDATE groups SET group_state=gstate, updated_at=now() WHERE id=gid;
  END
$$;

CREATE FUNCTION set_jobs_sync_v2(in_job_id bigint) RETURNS void
    LANGUAGE sql
    AS $$
  UPDATE jobs SET scheduler_sync = true, sync_count = sync_count-1 WHERE id = in_job_id;
$$;

CREATE FUNCTION sync_jobs_v2() RETURNS SETOF jobs
    LANGUAGE sql STABLE
    AS $$
  SELECT * FROM jobs WHERE (scheduler_sync = false) OR (sync_count > 0);
$$;

CREATE FUNCTION update_job_v3(p_job_id bigint, p_state text, p_build_started_at timestamp with time zone, p_build_finished_at timestamp with time zone, p_package_ident text, p_err_code integer, p_err_msg text) RETURNS void
    LANGUAGE sql
    AS $$
  UPDATE jobs
  SET job_state = p_state,
      scheduler_sync = false,
      sync_count = sync_count + 1,
      updated_at = now(),
      build_started_at = p_build_started_at,
      build_finished_at = p_build_finished_at,
      package_ident = p_package_ident,
      net_error_code = p_err_code,
      net_error_msg = p_err_msg
  WHERE id = p_job_id;
$$;

CREATE FUNCTION update_origin_package_vector_index() RETURNS trigger
    LANGUAGE plpgsql
    AS $$
    DECLARE iws TEXT;
    BEGIN
        NEW.ident_vector := to_tsvector(array_to_string(NEW.ident_array[1:2], ' '));
        RETURN NEW;
    END;
$$;

CREATE FUNCTION upsert_busy_worker_v1(in_ident text, in_job_id bigint, in_quarantined boolean) RETURNS SETOF busy_workers
    LANGUAGE plpgsql
    AS $$
  BEGIN
    RETURN QUERY INSERT INTO busy_workers (ident, job_id, quarantined)
    VALUES (in_ident, in_job_id, in_quarantined)
    ON CONFLICT(ident, job_id)
    DO UPDATE SET quarantined=in_quarantined RETURNING *;
    RETURN;
  END
$$;

ALTER TABLE ONLY origins_with_stats REPLICA IDENTITY NOTHING;

CREATE VIEW packages_with_channel_platform AS
 SELECT op.id,
    op.owner_id,
    op.name,
    op.ident,
    op.ident_array,
    op.checksum,
    op.manifest,
    op.config,
    op.target,
    op.deps,
    op.tdeps,
    op.exposes,
    op.visibility,
    op.created_at,
    op.updated_at,
    op.origin,
    array_agg(oc.name) OVER w AS channels,
    array_agg(op.target) OVER w AS platforms
   FROM ((origin_packages op
     JOIN origin_channel_packages ocp ON ((op.id = ocp.package_id)))
     JOIN origin_channels oc ON ((oc.id = ocp.channel_id)))
  WINDOW w AS (PARTITION BY op.ident);

ALTER TABLE ONLY account_tokens
    ADD CONSTRAINT account_tokens_account_id_key UNIQUE (account_id);

ALTER TABLE ONLY account_tokens
    ADD CONSTRAINT account_tokens_pkey PRIMARY KEY (id);

ALTER TABLE ONLY account_tokens
    ADD CONSTRAINT account_tokens_token_key UNIQUE (token);

ALTER TABLE ONLY accounts
    ADD CONSTRAINT accounts_name_key UNIQUE (name);

ALTER TABLE ONLY accounts
    ADD CONSTRAINT accounts_pkey PRIMARY KEY (id);

ALTER TABLE ONLY busy_workers
    ADD CONSTRAINT busy_workers_ident_job_id_key UNIQUE (ident, job_id);

ALTER TABLE ONLY group_projects
    ADD CONSTRAINT group_projects_pkey PRIMARY KEY (id);

ALTER TABLE ONLY groups
    ADD CONSTRAINT groups_pkey PRIMARY KEY (id);

ALTER TABLE ONLY jobs
    ADD CONSTRAINT jobs_pkey PRIMARY KEY (id);

ALTER TABLE ONLY origin_channel_packages
    ADD CONSTRAINT origin_channel_packages_pkey PRIMARY KEY (channel_id, package_id);

ALTER TABLE ONLY origin_channels
    ADD CONSTRAINT origin_channels_origin_key UNIQUE (origin, name);

ALTER TABLE ONLY origin_channels
    ADD CONSTRAINT origin_channels_pkey PRIMARY KEY (id);

ALTER TABLE ONLY origin_integrations
    ADD CONSTRAINT origin_integrations_origin_integration_name_key UNIQUE (origin, integration, name);

ALTER TABLE ONLY origin_integrations
    ADD CONSTRAINT origin_integrations_pkey PRIMARY KEY (id);

ALTER TABLE ONLY origin_invitations
    ADD CONSTRAINT origin_invitations_pkey PRIMARY KEY (id);

ALTER TABLE ONLY origin_packages
    ADD CONSTRAINT origin_packages_ident_key UNIQUE (ident);

ALTER TABLE ONLY origin_packages
    ADD CONSTRAINT origin_packages_pkey PRIMARY KEY (id);

ALTER TABLE ONLY origin_private_encryption_keys
    ADD CONSTRAINT origin_private_encryption_keys_full_name_key UNIQUE (full_name);

ALTER TABLE ONLY origin_private_encryption_keys
    ADD CONSTRAINT origin_private_encryption_keys_pkey PRIMARY KEY (id);

ALTER TABLE ONLY origin_project_integrations
    ADD CONSTRAINT origin_project_integrations_pkey PRIMARY KEY (id);

ALTER TABLE ONLY origin_project_integrations
    ADD CONSTRAINT origin_project_integrations_project_id_integration_id_key UNIQUE (project_id, integration_id);

ALTER TABLE ONLY origin_projects
    ADD CONSTRAINT origin_projects_origin_package_name_name_key UNIQUE (origin, package_name, name);

ALTER TABLE ONLY origin_projects
    ADD CONSTRAINT origin_projects_pkey PRIMARY KEY (id);

ALTER TABLE ONLY origin_public_encryption_keys
    ADD CONSTRAINT origin_public_encryption_keys_full_name_key UNIQUE (full_name);

ALTER TABLE ONLY origin_public_encryption_keys
    ADD CONSTRAINT origin_public_encryption_keys_pkey PRIMARY KEY (id);

ALTER TABLE ONLY origin_public_keys
    ADD CONSTRAINT origin_public_keys_full_name_key UNIQUE (full_name);

ALTER TABLE ONLY origin_public_keys
    ADD CONSTRAINT origin_public_keys_pkey PRIMARY KEY (id);

ALTER TABLE ONLY origin_secret_keys
    ADD CONSTRAINT origin_secret_keys_full_name_key UNIQUE (full_name);

ALTER TABLE ONLY origin_secret_keys
    ADD CONSTRAINT origin_secret_keys_pkey PRIMARY KEY (id);

ALTER TABLE ONLY origin_secrets
    ADD CONSTRAINT origin_secrets_pkey PRIMARY KEY (id);

ALTER TABLE ONLY origins
    ADD CONSTRAINT origins_name_key UNIQUE (name);

ALTER TABLE ONLY origins
    ADD CONSTRAINT origins_pkey PRIMARY KEY (name);

CREATE INDEX ident_index ON origin_packages USING gin (ident_vector);

CREATE INDEX origin_packages_ident_array ON origin_packages USING btree (ident_array);

CREATE INDEX pending_groups_index_v1 ON groups USING btree (created_at) WHERE (group_state = 'Pending'::text);

CREATE INDEX pending_jobs_index_v1 ON jobs USING btree (created_at) WHERE (job_state = 'Pending'::text);

CREATE INDEX queued_groups_index_v1 ON groups USING btree (created_at) WHERE (group_state = 'Queued'::text);

CREATE RULE "_RETURN" AS
    ON SELECT TO origins_with_stats DO INSTEAD  SELECT o.name,
    o.owner_id,
    o.created_at,
    o.updated_at,
    o.default_package_visibility,
    count(DISTINCT op.name) AS package_count
   FROM (origins o
     LEFT JOIN origin_packages op ON ((o.name = op.origin)))
  GROUP BY o.name
  ORDER BY o.name DESC;

CREATE TRIGGER origin_packages_vector BEFORE INSERT ON origin_packages FOR EACH ROW EXECUTE PROCEDURE update_origin_package_vector_index();

ALTER TABLE ONLY audit_package_group
    ADD CONSTRAINT audit_package_group_origin_fkey FOREIGN KEY (origin) REFERENCES origins(name);

ALTER TABLE ONLY audit_package
    ADD CONSTRAINT audit_package_origin_fkey FOREIGN KEY (origin) REFERENCES origins(name);

ALTER TABLE ONLY origin_channel_packages
    ADD CONSTRAINT origin_channel_packages_channel_id_fkey FOREIGN KEY (channel_id) REFERENCES origin_channels(id) ON DELETE CASCADE;

ALTER TABLE ONLY origin_channel_packages
    ADD CONSTRAINT origin_channel_packages_package_id_fkey FOREIGN KEY (package_id) REFERENCES origin_packages(id);

ALTER TABLE ONLY origin_channels
    ADD CONSTRAINT origin_channels_origin_fkey FOREIGN KEY (origin) REFERENCES origins(name);

ALTER TABLE ONLY origin_invitations
    ADD CONSTRAINT origin_invitations_origin_fkey FOREIGN KEY (origin) REFERENCES origins(name);

ALTER TABLE ONLY origin_members
    ADD CONSTRAINT origin_members_origin_fkey FOREIGN KEY (origin) REFERENCES origins(name);

ALTER TABLE ONLY origin_packages
    ADD CONSTRAINT origin_packages_origin_fkey FOREIGN KEY (origin) REFERENCES origins(name);

ALTER TABLE ONLY origin_private_encryption_keys
    ADD CONSTRAINT origin_private_encryption_keys_origin_fkey FOREIGN KEY (origin) REFERENCES origins(name);

ALTER TABLE ONLY origin_project_integrations
    ADD CONSTRAINT origin_project_integrations_integration_id_fkey FOREIGN KEY (integration_id) REFERENCES origin_integrations(id) ON DELETE CASCADE;

ALTER TABLE ONLY origin_project_integrations
    ADD CONSTRAINT origin_project_integrations_project_id_fkey FOREIGN KEY (project_id) REFERENCES origin_projects(id) ON DELETE CASCADE;

ALTER TABLE ONLY origin_projects
    ADD CONSTRAINT origin_projects_origin_fkey FOREIGN KEY (origin) REFERENCES origins(name);

ALTER TABLE ONLY origin_public_encryption_keys
    ADD CONSTRAINT origin_public_encryption_keys_origin_fkey FOREIGN KEY (origin) REFERENCES origins(name);

ALTER TABLE ONLY origin_public_keys
    ADD CONSTRAINT origin_public_keys_origin_fkey FOREIGN KEY (origin) REFERENCES origins(name);

ALTER TABLE ONLY origin_secret_keys
    ADD CONSTRAINT origin_secret_keys_origin_fkey FOREIGN KEY (origin) REFERENCES origins(name);

ALTER TABLE ONLY origin_secrets
    ADD CONSTRAINT origin_secrets_origin_fkey FOREIGN KEY (origin) REFERENCES origins(name);
