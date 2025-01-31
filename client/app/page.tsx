"use client";
import { useState, useEffect } from "react";
import { useRouter } from "next/navigation";
import { Button, Card, Container, Title, Table, Pagination } from "@mantine/core";
import '@mantine/core/styles.css';
import Cookies from "js-cookie";
import ErrorNotification from "@/app/components/ErrorNotification";

export default function MainPage() {
    const [orders, setOrders] = useState([]);
    const [currentPage, setCurrentPage] = useState(1);
    const [totalOrders, setTotalOrders] = useState(0);
    const [errorMessage, setErrorMessage] = useState<string | null>(null);
    const router = useRouter();

    const itemsPerPage = 10;

    useEffect(() => {
        const jwtToken = Cookies.get("jwt");
        const refreshToken = Cookies.get("refresh");
        if (jwtToken) {
            fetchOrders(jwtToken, currentPage, itemsPerPage);
        } else if (!jwtToken && refreshToken) {
            handleRefresh();
        } else {
            router.push("/sign-in");
        }
    }, [currentPage, router]);

    const handleRefresh = async () => {
        const response = await fetch("http://localhost:8081/user/refresh", {
            method: "POST",
            headers: { "Content-Type": "application/json" },
            credentials: 'include',
        });
        if (response.ok) {
            const data = await response.json();
            Cookies.set("jwt", data.jwt); // Assuming the refreshed token is in the response
            fetchOrders(data.jwt, currentPage, itemsPerPage);
        }
    };

    const fetchOrders = async (token: string, page: number, limit: number) => {
        const response = await fetch(`http://localhost:8082/api/orders?page=${page}&limit=${limit}`, {
            method: "GET",
            headers: {
                "Content-Type": "application/json",
                "Authorization": `Bearer ${token}`,
            },
            credentials: 'include',
        });

        if (response.ok) {
            const data = await response.json();
            setOrders(data.orders);
            setTotalOrders(data.total);
        } else {
            const responseMessage = await response.json();
            setErrorMessage(responseMessage.error || "Authentication failed");
            setTimeout(() => setErrorMessage(null), 7000);
        }
    };

    const handleSignOut = async () => {
        Cookies.remove("refresh");
        Cookies.remove("jwt");
        router.push("/sign-in");
    };

    return (
        <>
            <ErrorNotification errorMessage={errorMessage} onClose={() => setErrorMessage(null)} />

            <Container size="sm" mt={50}>
                <Card shadow="sm" p="lg" radius="md" withBorder mt="xl">
                    <Title order={3}>Your Orders</Title>
                    <Table mt="md">
                        <Table.Thead>
                            <Table.Tr>
                                <Table.Th>Order ID</Table.Th>
                                <Table.Th>Product Name</Table.Th>
                                <Table.Th>Client Name</Table.Th>
                                <Table.Th>Price</Table.Th>
                            </Table.Tr>
                        </Table.Thead>
                        <Table.Tbody>
                            {orders.map((order: any) => (
                                <Table.Tr key={order.id}>
                                    <Table.Td>{order.id}</Table.Td>
                                    <Table.Td>{order.productName}</Table.Td>
                                    <Table.Td>{order.clientName}</Table.Td>
                                    <Table.Td>{order.price}</Table.Td>
                                </Table.Tr>
                            ))}
                        </Table.Tbody>
                    </Table>
                    <Pagination
                        page={currentPage}
                        onChange={(page) => {
                            setCurrentPage(page);
                            fetchOrders(Cookies.get("jwt") || "", page, itemsPerPage); // Fetch orders for the new page
                        }}
                        total={Math.ceil(totalOrders / itemsPerPage)}
                        mt="md"
                    />
                </Card>

                <Card shadow="sm" p="lg" radius="md" withBorder>
                    <Button variant="outline" color="red" size="xl" radius="xl" onClick={handleSignOut}>
                        Logout
                    </Button>
                </Card>
            </Container>
        </>
    );
}
