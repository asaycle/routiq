import React, { useEffect, useState } from 'react';
import {
    Autocomplete,
    Chip,
    Container,
    List,
    ListItem,
    ListItemButton,
    ListItemText,
    Paper,
    Typography,
    Box,
    TextField,
    Button,
    MenuItem,
} from '@mui/material';
import { useNavigate, useParams } from 'react-router-dom';
import { LocalizationProvider } from '@mui/x-date-pickers';
import { AdapterDateFns } from '@mui/x-date-pickers/AdapterDateFns';
import { DatePicker } from '@mui/x-date-pickers';
import { createRiding } from '../../api/ridingClient';
import { createTag, listTags } from '../../api/tagClient';
import { Tag } from '../../../lib/proto/v1/tag_pb';

const RidingCreate: React.FC = () => {
    const navigate = useNavigate();
    const { id } = useParams<{ id: string }>();
    const validID: string | null = id ?? null;
    const [date, setDate] = useState<Date | null>(new Date());
    const [title, setTitle] = useState('');
    const [note, setNote] = useState('');
    const [score, setScore] = useState(0);
    const [isSubmitting, setIsSubmitting] = useState<boolean>(false);
    const [filter, setFilter] = useState("");
    const [tags, setTags] = useState<Tag[]>([]); // タグ候補
    const [selectedTags, setSelectedTags] = useState<Tag[]>([]); // 選択されたタグ
    const [isCreating, setIsCreating] = useState(false);


    const handleSubmit = async (e: React.FormEvent) => {
        e.preventDefault();

        // サーバーへ送信する処理を追加
        try {
            const riding = await createRiding(validID, title, date, score, tags);
            alert('ルートが作成されました');
            console.log('Created Route:', riding);
            navigate(`/routes/${id}`);
        } catch (error) {
            console.error('Error creating route:', error);
            alert('ルートの作成中にエラーが発生しました');
        } finally {
            setIsSubmitting(false);
        }
    };

    useEffect(() => {
        listTags("").then((tags) => setTags(tags)).catch(console.error);
    }, [])

    useEffect(() => {
        if (filter) {
            listTags(filter).then((tags) => setTags(tags)).catch(console.error);
        }
    }, [filter])

    const handleSelectTag = (tag: Tag) => {
        if (!selectedTags.find((t) => t.getId() === tag.getId())) {
            setSelectedTags([...selectedTags, tag]);
        }
    };

    // タグ選択解除
    const handleRemoveTag = (tagId: string) => {
        setSelectedTags(selectedTags.filter((tag) => tag.getId() !== tagId));
    };

    const handleCreateTag = async () => {
        setIsCreating(true);
        try {
            const tag = await createTag(filter);
            setSelectedTags([...selectedTags, tag]);
            setTags([...tags, tag]); // 新しいタグを候補に追加
            setFilter(""); // 入力をリセット
        } catch (error) {
            console.error("Failed to create new tag:", error);
        } finally {
            setIsCreating(false);
        }
    };

    return (
        <LocalizationProvider dateAdapter={AdapterDateFns}>
            <Container maxWidth="md" style={{ marginTop: '20px' }}>
                <Paper elevation={3} style={{ padding: '20px' }}>
                    <Typography variant="h5" gutterBottom>
                        ツーリング記録を作成
                    </Typography>
                    <form onSubmit={handleSubmit}>
                        {/* 走行日 */}
                        <Box mt={2}>
                            <DatePicker
                                label="走行日"
                                value={date}
                                onChange={(newValue) => setDate(newValue)}
                                slots={{
                                    textField: (props) => <TextField {...props} fullWidth />,
                                }}
                            />
                        </Box>

                        {/* タイトル */}
                        <Box mt={2}>
                            <TextField
                                label="タイトル"
                                value={title}
                                onChange={(e) => setTitle(e.target.value)}
                                variant="outlined"
                                fullWidth
                                required
                            />
                        </Box>

                        {/* メモ */}
                        <Box mt={2}>
                            <TextField
                                label="メモ"
                                value={note}
                                onChange={(e) => setNote(e.target.value)}
                                variant="outlined"
                                multiline
                                rows={4}
                                fullWidth
                            />
                        </Box>

                        {/* スコア */}
                        <Box mt={2}>
                            <TextField
                                select
                                label="スコア"
                                value={score}
                                onChange={(e) => setScore(Number(e.target.value))}
                                variant="outlined"
                                fullWidth
                            >
                                {[0, 1, 2, 3, 4, 5].map((option) => (
                                    <MenuItem key={option} value={option}>
                                        {option}
                                    </MenuItem>
                                ))}
                            </TextField>
                        </Box>

                        <Box mt={2}>
                            <Autocomplete
                                multiple
                                options={tags}
                                getOptionLabel={(option) => option.getName()}
                                value={selectedTags}
                                onChange={(event, newValue) => {
                                    setSelectedTags(newValue);
                                }}
                                renderTags={(value, getTagProps) =>
                                    value.map((option, index) => (
                                        <Chip
                                            label={option.getName()}
                                            {...getTagProps({ index })}
                                            key={option.getId()}
                                        />
                                    ))
                                }
                                renderInput={(params) => (
                                    <TextField {...params} label="Tags" placeholder="Select tags" />
                                )}
                                onInputChange={(event, value) => {
                                    setFilter(value);
                                    listTags(value).then((tags) => setTags(tags)).catch(console.error);
                                }}
                                noOptionsText={
                                    filter ? (
                                        <ListItemButton onClick={handleCreateTag} disabled={isCreating}>
                                            <ListItemText primary={`+ タグを新規作成します: "${filter}"`} />
                                        </ListItemButton>
                                    ) : (
                                        "フィルタが存在しません"
                                    )
                                }
                            />
                            {filter && !tags.some((tag) => tag.getName() === filter) && (
                                <List>
                                    <ListItem disablePadding>
                                        <ListItemButton
                                            onClick={handleCreateTag}
                                            disabled={isCreating}
                                        >
                                            <ListItemText primary={`+ Create new tag: "${filter}"`} />
                                        </ListItemButton>
                                    </ListItem>
                                </List>
                            )}
                        </Box>

                        {/* ボタン */}
                        <Box mt={4} display="flex" justifyContent="space-between">
                            <Button variant="outlined" color="secondary" onClick={() => console.log('キャンセル')}>
                                キャンセル
                            </Button>
                            <Button type="submit" variant="contained" color="primary">
                                作成
                            </Button>
                        </Box>
                    </form>
                </Paper>
            </Container>
        </LocalizationProvider>
    );
};

export default RidingCreate;
