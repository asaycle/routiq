// src/pages/AnnouncementsPage.tsx
import React from "react";
import { Box, Pagination, Stack, Typography } from "@mui/material";
import { listAnnouncements } from "../../api/announcementClient";
import AnnouncementList from "../../components/announcement/AnnouncementList";
import { Announcement } from '../../../lib/proto/v1/announcement_pb';


const PAGE_SIZE = 10;


export default function AnnouncementsPage() {
    const [page, setPage] = React.useState(1);
    const [announcements, setAnnouncements] = React.useState<Announcement[]>([]);
    const [total, setTotal] = React.useState(0);
    const [loading, setLoading] = React.useState(true);
    const [error, setError] = React.useState<string | null>(null);


    const fetchPage = React.useCallback(async (p: number) => {
        try {
            setLoading(true);
            setAnnouncements(await listAnnouncements("", PAGE_SIZE));
        } catch (e) {
            console.error("Error fetching announcements:", e);
            setError("アナウンスの取得に失敗しました");
        } finally {
            setLoading(false);
        }
    }, []);


    React.useEffect(() => {
        fetchPage(page);
    }, [fetchPage, page]);


    const pageCount = Math.max(1, Math.ceil(total / PAGE_SIZE));


    return (
        <Stack spacing={5}>
            <Typography variant="h4" component="h1">お知らせ一覧</Typography>
            <AnnouncementList items={announcements} loading={loading} error={error} itemHref={(a) => `/${a.getName()}`} />
            <Box display="flex" justifyContent="center" sx={{ mt: 1 }}>
                <Pagination
                    page={page}
                    count={pageCount}
                    onChange={(_, p) => setPage(p)}
                    color="primary"
                    shape="rounded"
                    showFirstButton
                    showLastButton
                />
            </Box>
        </Stack>
    );
}