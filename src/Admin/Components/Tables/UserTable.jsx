import React, { useState } from 'react'
import {
    Table,
    TableBody,
    TableCell,
    TableContainer,
    TableHead,
    TableRow,
    Paper,
    Container,
    InputAdornment,
    TextField,
    Chip
} from '@mui/material'
import { Link } from 'react-router-dom';
import { AiOutlineSearch } from 'react-icons/ai';
import AddUser from '../AddUser';

const UserTable = ({ user, getUser }) => {
    const columns = [
        {
            id: 'name',
            label: 'Name',
            minWidth: 100,
            align: 'center',
        },
        {
            id: 'email',
            label: 'Email',
            minWidth: 150,
            align: 'center',
        },
        {
            id: 'role',
            label: 'Role',
            minWidth: 100,
            align: 'center',
        },
        {
            id: 'phone',
            label: 'Phone',
            align: 'center',
            minWidth: 120
        },
        {
            id: 'date',
            label: 'Created On',
            minWidth: 170,
            align: 'center',
        },
    ];

    const [searchQuery, setSearchQuery] = useState("");
    
    const handleSearchInputChange = (event) => {
        setSearchQuery(event.target.value);
    };

    // Sort users by creation date (newest first)
    const sortedUsers = [...user].sort((a, b) => new Date(b.createdAt) - new Date(a.createdAt));

    const filteredUsers = sortedUsers.filter((user) => {
        const searchLower = searchQuery.toLowerCase();
        return (
            user.firstName?.toLowerCase().includes(searchLower) ||
            user.lastName?.toLowerCase().includes(searchLower) ||
            user.email?.toLowerCase().includes(searchLower) ||
            user.phone?.toLowerCase().includes(searchLower) ||
            user.role?.toLowerCase().includes(searchLower)
        );
    });

    const getRoleColor = (role) => {
        switch(role?.toLowerCase()) {
            case 'admin':
                return 'error';
            case 'user':
                return 'primary';
            default:
                return 'default';
        }
    };

    return (
        <>
            <Container sx={{ display: 'flex', justifyContent: 'center', alignItems: 'center', marginBottom: 5, marginTop: 5 }}>
                <TextField
                    id="search"
                    type="search"
                    label="Search Users"
                    onChange={handleSearchInputChange}
                    className="placeholder-animation"
                    sx={{ width: { xs: 350, sm: 500, md: 800 }, }}
                    InputProps={{
                        endAdornment: (
                            <InputAdornment position="end">
                                <AiOutlineSearch />
                            </InputAdornment>
                        ),
                    }}
                />
            </Container>
            
            <AddUser getUser={getUser} user={sortedUsers} />
            
            <Paper style={{overflow: "auto"}}>
                <TableContainer sx={{ maxHeight: '400px' }}>
                    <Table stickyHeader aria-label="sticky table">
                        <TableHead>
                            <TableRow>
                                {columns.map((column) => (
                                    <TableCell
                                        key={column.id}
                                        align={column.align}
                                        style={{ 
                                            minWidth: column.minWidth, 
                                            color: "#1976d2",
                                            fontWeight: 'bold',
                                            backgroundColor: '#fff'
                                        }}
                                    >
                                        {column.label}
                                    </TableCell>
                                ))}
                            </TableRow>
                        </TableHead>
                        <TableBody>
                            {filteredUsers.length === 0 ? (
                                <TableRow>
                                    <TableCell colSpan={columns.length}>
                                        <div style={{ display: "flex", justifyContent: "center" }}>
                                            <h4>No users found.</h4>
                                        </div>
                                    </TableCell>
                                </TableRow>
                            ) : (
                                filteredUsers.map((user) => (
                                    <TableRow
                                        key={user.id}
                                        hover
                                        role="checkbox"
                                        tabIndex={-1}
                                    >
                                        <TableCell align="center">
                                            <Link to={`user/${user.id}`}>
                                                {user.firstName} {user.lastName}
                                            </Link>
                                        </TableCell>
                                        <TableCell align="center">
                                            <Link to={`user/${user.id}`}>
                                                {user.email}
                                            </Link>
                                        </TableCell>
                                        <TableCell align="center">
                                            <Chip 
                                                label={user.role}
                                                color={getRoleColor(user.role)}
                                                size="small"
                                            />
                                        </TableCell>
                                        <TableCell align="center">
                                            <Link to={`user/${user.id}`}>
                                                {user.phone || 'N/A'}
                                            </Link>
                                        </TableCell>
                                        <TableCell align="center">
                                            <Link to={`user/${user.id}`}>
                                                {new Date(user.createdAt).toLocaleDateString('en-US', {
                                                    weekday: "long",
                                                    year: "numeric",
                                                    month: "short",
                                                    day: "numeric"
                                                })}
                                                {" "}
                                                {new Date(user.createdAt).toLocaleTimeString('en-US')}
                                            </Link>
                                        </TableCell>
                                    </TableRow>
                                ))
                            )}
                        </TableBody>
                    </Table>
                </TableContainer>
            </Paper>
        </>
    );
};

export default UserTable;