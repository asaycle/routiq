import React from "react";
import { Card, CardActionArea, CardContent, Stack, Typography } from "@mui/material";
import moment from "moment";
import { Announcement } from '../../../lib/proto/v1/announcement_pb';


export type AnnouncementItemProps = {
    announcement: Announcement;
    onClick?: () => void;
    href?: string;
};

const truncate = (text: string, len = 120) =>
    text.length > len ? `${text.slice(0, len)}…` : text;

function parseProtoTimestamp(ts?: { seconds?: number; nanos?: number }) {
    console.log("TS", ts?.seconds);
    if (!ts || ts.seconds === undefined) return null;
    const millis = ts.seconds * 1000 + Math.floor((ts.nanos ?? 0) / 1e6);
    return moment(millis); // momentオブジェクトが返る
}

export default function AnnouncementItem({
    announcement,
    onClick,
    href,
}: AnnouncementItemProps) {
    const content = (
        <CardActionArea onClick={onClick} disableRipple={!href && !onClick}>
            <CardContent>
                <Stack spacing={1}>
                    <Typography variant="h6" component="h3" sx={{ lineHeight: 1.3 }}>
                        {announcement.getTitle() || "No Title"}
                    </Typography>
                    <Stack direction="row" spacing={1} alignItems="center">
                        <Typography variant="caption" color="text.secondary">
                            {parseProtoTimestamp(announcement.getCreateTime())?.format("YYYY/MM/DD HH:mm")}
                        </Typography>
                    </Stack>
                    <Typography variant="body2" color="text.secondary">
                        {truncate(announcement.getBody())}
                    </Typography>
                </Stack>
            </CardContent>
        </CardActionArea>
    );

    return <Card variant="outlined">{content}</Card>;
}
