// src/components/announcements/AnnouncementSection.tsx
import React from "react";
import { Box, Link as MuiLink, Stack, Typography } from "@mui/material";
import { listAnnouncements } from "../../api/announcementClient";
import { Announcement } from '../../../lib/proto/v1/announcement_pb';
import AnnouncementList from "./AnnouncementList";
import { Link as RouterLink } from "react-router-dom";


export default function AnnouncementSection() {
    const [items, setItems] = React.useState<Announcement[]>([]);
    const [loading, setLoading] = React.useState(true);
    const [error, setError] = React.useState<string | null>(null);


    React.useEffect(() => {
        (async () => {
            try {
                setLoading(true);
                const latest = await listAnnouncements("", 5);
                setItems(latest);
                setError(null);
            } catch (e) {
                setError("最新のアナウンスを取得できませんでした");
            } finally {
                setLoading(false);
            }
        })();
    }, []);


    return (
        <Stack spacing={1.5}>
            <Box display="flex" alignItems="baseline" justifyContent="space-between">
                <Typography variant="h5" component="h2">最新のお知らせ</Typography>
                <MuiLink component={RouterLink} to="/announcements" underline="hover">
                    もっと見る
                </MuiLink>
            </Box>
            <AnnouncementList items={items} loading={loading} error={error} dense itemHref={(a) => `/${a.getName()}`} />
        </Stack>
    );
}