CREATE TABLE access_rules (
  id varchar primary key,
  ptype varchar,
  v0 varchar,
  v1 varchar,
  v2 varchar,
  v3 varchar,
  v4 varchar,
  v5 varchar,
  deleted_at timestamp
);

CREATE INDEX access_rule_ptype ON access_rules (ptype);
CREATE INDEX access_rule_v0 ON access_rules (v0);
CREATE INDEX access_rule_v1 ON access_rules (v1);
CREATE INDEX access_rule_v2 ON access_rules (v2);
CREATE INDEX access_rule_v3 ON access_rules (v3);
CREATE INDEX access_rule_v4 ON access_rules (v4);
CREATE INDEX access_rule_v5 ON access_rules (v5);
CREATE INDEX access_rule_deleted_at ON access_rules (deleted_at);
