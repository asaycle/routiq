import React, { useEffect, useState } from 'react';
import { listTags } from '../../api/tagClient';
import { Table, TableBody, TableCell, TableContainer, TableHead, TableRow, Paper, Button } from '@mui/material';

const TagList: React.FC = () => {
  const [tags, setTags] = useState<string[]>([]);

  useEffect(() => {
    const fetchTags = async () => {
      try {
        const tagList = await listTags("");
        setTags(tagList.map(tag => tag.getName()));
      } catch (error) {
        console.error('Error fetching tags:', error);
      }
    };

    fetchTags();
  }, []);


  return (
    <div>
      <TableContainer component={Paper}>
        <Table>
          <TableHead>
            <TableRow>
              <TableCell>ID</TableCell>
              <TableCell>Name</TableCell>
              <TableCell>Action</TableCell>
            </TableRow>
          </TableHead>
          <TableBody>
            {tags.map((tag, index) => (
              <TableRow key={index}>
                <TableCell>{index}</TableCell>
                <TableCell>{tag}</TableCell>
                <TableCell>
                  <Button variant="contained" color="primary">
                    View
                  </Button>
                </TableCell>
              </TableRow>
            ))}
          </TableBody>
        </Table>
      </TableContainer>
    </div>
  );
};

export default TagList;
