import React from 'react';
import { Typography, Container } from '@mui/material';
import { listAnnouncements } from "../api/announcementClient";
import AnnouncementList from "../components/announcement/AnnouncementList";
import { Announcement } from '../../lib/proto/v1/announcement_pb';

const PAGE_SIZE = 10;

const HomePage: React.FC = () => {
    const [page, setPage] = React.useState(1);
    const [announcements, setAnnouncements] = React.useState<Announcement[]>([]);
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

    return (
        <Container>
            <Typography variant="h4" gutterBottom>
                Routiqへようこそ
            </Typography>
            <Typography variant="h4" component="h1">お知らせ一覧</Typography>
            <AnnouncementList items={announcements} loading={loading} error={error} itemHref={(a) => `/${a.getName()}`} />
        </Container>
    );
};

export default HomePage;
