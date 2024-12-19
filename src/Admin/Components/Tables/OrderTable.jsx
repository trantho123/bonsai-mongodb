import React, { useState } from 'react';
import {
    Table,
    TableBody,
    TableCell,
    TableContainer,
    TableHead,
    TableRow,
    Paper,
    IconButton,
    Collapse,
    Typography,
    Box,
    Chip
} from '@mui/material';
import { MdKeyboardArrowDown } from 'react-icons/md'

const OrderTable = ({ orders }) => {
    const [openOrderId, setOpenOrderId] = useState("");
    
    // Sort orders by date, most recent first
    const sortedOrders = orders?.sort((a, b) => 
        new Date(b.createdAt) - new Date(a.createdAt)
    ) || [];

    const formatCurrency = (amount) => {
        if (!amount && amount !== 0) return 'N/A';
        return new Intl.NumberFormat('vi-VN', { 
            style: 'currency', 
            currency: 'VND' 
        }).format(amount);
    };

    const getStatusColor = (status) => {
        const colors = {
            pending: 'warning',
            confirmed: 'info',
            completed: 'success',
            failed: 'error',
            default: 'default'
        };
        return colors[status?.toLowerCase()] || colors.default;
    };

    if (!orders || orders.length === 0) {
        return (
            <Box sx={{ textAlign: 'center', p: 3 }}>
                <Typography variant="h6">No orders available</Typography>
            </Box>
        );
    }

    return (
        <Paper style={{ overflow: "auto", maxHeight: "500px" }}>
            <TableContainer sx={{ maxHeight: '500px' }}>
                <Table stickyHeader aria-label="sticky table">
                    <TableHead>
                        <TableRow>
                            <TableCell />
                            <TableCell sx={{ color: "#1976d2", fontWeight: 'bold' }}>Order ID</TableCell>
                            <TableCell sx={{ color: "#1976d2", fontWeight: 'bold' }}>User ID</TableCell>
                            <TableCell sx={{ color: "#1976d2", fontWeight: 'bold' }}>Status</TableCell>
                            <TableCell sx={{ color: "#1976d2", fontWeight: 'bold' }}>Amount</TableCell>
                            <TableCell sx={{ color: "#1976d2", fontWeight: 'bold' }}>Order Date</TableCell>
                        </TableRow>
                    </TableHead>
                    <TableBody>
                        {sortedOrders.map((order) => (
                            <React.Fragment key={order.id}>
                                <TableRow hover>
                                    <TableCell>
                                        <IconButton
                                            aria-label="expand row"
                                            size="small"
                                            onClick={() => setOpenOrderId(openOrderId === order.id ? "" : order.id)}
                                        >
                                            <MdKeyboardArrowDown 
                                                style={{ 
                                                    transform: openOrderId === order.id ? 'rotate(180deg)' : 'rotate(0)',
                                                    transition: 'transform 0.2s'
                                                }}
                                            />
                                        </IconButton>
                                    </TableCell>
                                    <TableCell>{order.id}</TableCell>
                                    <TableCell>{order.user}</TableCell>
                                    <TableCell>
                                        <Chip 
                                            label={order.status || 'N/A'}
                                            color={getStatusColor(order.status)}
                                            size="small"
                                        />
                                    </TableCell>
                                    <TableCell>{formatCurrency(order.amount)}</TableCell>
                                    <TableCell>
                                        {new Date(order.createdAt).toLocaleDateString('en-us', {
                                            year: "numeric", 
                                            month: "short", 
                                            day: "numeric"
                                        })}
                                        {" "}
                                        {new Date(order.createdAt).toLocaleTimeString('en-US')}
                                    </TableCell>
                                </TableRow>
                                <TableRow>
                                    <TableCell style={{ paddingBottom: 0, paddingTop: 0 }} colSpan={6}>
                                        <Collapse in={openOrderId === order.id} timeout="auto" unmountOnExit>
                                            <Box sx={{ margin: 2 }}>
                                                <Typography variant="h6" gutterBottom component="div">
                                                    Order Details
                                                </Typography>
                                                <Box sx={{ mb: 2 }}>
                                                    <Typography>{`Order ID: ${order.id}`}</Typography>
                                                    <Typography>{`User ID: ${order.user}`}</Typography>
                                                    <Typography>{`Amount: ${formatCurrency(order.amount)}`}</Typography>
                                                    <Typography>{`Status: ${order.status}`}</Typography>
                                                    <Typography>{`Created At: ${new Date(order.createdAt).toLocaleString()}`}</Typography>
                                                </Box>
                                            </Box>
                                        </Collapse>
                                    </TableCell>
                                </TableRow>
                            </React.Fragment>
                        ))}
                    </TableBody>
                </Table>
            </TableContainer>
        </Paper>
    );
};

export default OrderTable;
