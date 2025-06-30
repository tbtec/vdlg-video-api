create table if not exists vdlg_video.video
(
    video_id uuid not null primary key,
    customer_id uuid, 
    status varchar(200) not null,
    file_name_input varchar(200) not null,
    file_name_output varchar(200) not null,
    created_at timestamp,
    updated_at timestamp
);
