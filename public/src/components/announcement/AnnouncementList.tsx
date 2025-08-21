// src/components/announcements/AnnouncementList.tsx
import React from "react";
import { Alert, Box, Skeleton, Stack } from "@mui/material";
import AnnouncementItem from "./AnnouncementItem";
import { Announcement } from '../../../lib/proto/v1/announcement_pb';

export type AnnouncementListProps = {
    items: Announcement[];
    loading?: boolean;
    error?: string | null;
    onItemClick?: (a: Announcement) => void;
    itemHref?: (a: Announcement) => string | undefined;
    dense?: boolean; // Topなどで余白を詰めたい場合にtrue
};

export default function AnnouncementList({
    items,
    loading,
    error,
    onItemClick,
    itemHref,
    dense,
}: AnnouncementListProps) {
    if (loading) {
        return (
            <Stack spacing={dense ? 1 : 2}>
                {Array.from({ length: 5 }).map((_, i) => (
                    <Skeleton key={i} variant="rounded" height={88} />
                ))}
            </Stack>
        );
    }

    if (error) {
        return <Alert severity="error">{error}</Alert>;
    }

    if (!items || items.length === 0) {
        return <Alert severity="info">アナウンスはまだありません。</Alert>;
    }

    return (
        <Stack spacing={dense ? 1 : 2}>
            {items.map((a) => (
                <Box key={a.getName()}>
                    <AnnouncementItem
                        announcement={a}
                        onClick={onItemClick ? () => onItemClick(a) : undefined}
                        href={itemHref ? itemHref(a) : undefined}
                    />
                </Box>
            ))}
        </Stack>
    );
}
