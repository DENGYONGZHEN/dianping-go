CREATE TABLE tb_user (
    id BIGSERIAL PRIMARY KEY,
    phone VARCHAR(11) NOT NULL,
    password VARCHAR(128) DEFAULT '',
    nick_name VARCHAR(32) DEFAULT '',
    icon VARCHAR(255) DEFAULT '',
    create_time TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    update_time TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE UNIQUE INDEX uniqe_key_phone ON tb_user (phone);

ALTER SEQUENCE tb_user_id_seq RESTART WITH 1010;

COMMENT ON TABLE tb_user IS '用户表';
COMMENT ON COLUMN tb_user.id IS '主键';
COMMENT ON COLUMN tb_user.phone IS '手机号码';
COMMENT ON COLUMN tb_user.password IS '密码，加密存储';
COMMENT ON COLUMN tb_user.nick_name IS '昵称，默认是用户id';
COMMENT ON COLUMN tb_user.icon IS '人物头像';
COMMENT ON COLUMN tb_user.create_time IS '创建时间';
COMMENT ON COLUMN tb_user.update_time IS '更新时间';


