"use client";
import { useState, useEffect } from "react";
import { useRouter } from "next/navigation";
import { Button, Card, Text, Container, Title, Space, Table, Pagination } from "@mantine/core";
import '@mantine/core/styles.css';
import Cookies from "js-cookie";

export default function MainPage() {
    const [orders, setOrders] = useState([]);
    const [currentPage, setCurrentPage] = useState(1);
    const [totalOrders, setTotalOrders] = useState(0);
    const [token, setToken] = useState<string | null>(null);
    const router = useRouter();

    useEffect(() => {
        const storedToken = Cookies.get("jwt");
        if (storedToken) {
            setToken(storedToken);
            fetchOrders(storedToken, currentPage);
        }
    }, [currentPage]);

    const fetchOrders = async (token: string, page: number) => {
        const response = await fetch(`http://localhost:8081/orders?page=${page}`, {
            method: "GET",
            headers: {
                "Content-Type": "application/json",
                "Authorization": `Bearer ${token}`,
            },
        });

        if (response.ok) {
            const data = await response.json();
            setOrders(data.orders); // assuming the response has orders array
            setTotalOrders(data.total); // assuming the response has the total number of orders
        }
    };

    const handleSignInRedirect = () => {
        router.push("/sign-in"); // Redirect to the sign-in page if the user is not logged in
    };

    if (!token) {
        return (
            <Container size="sm" mt={50}>
                <Card shadow="sm" p="lg" radius="md" withBorder>
                    <Title align="center" order={2} mb="lg">Your session has expired. Please login to be able to view all orders</Title>
                    <Button variant="filled" color="blue" size="xl" radius="xl" onClick={handleSignInRedirect}>
                        Login
                    </Button>
                </Card>
            </Container>
        );
    }

    return (
        <Container size="sm" mt={50}>
            <Card shadow="sm" p="lg" radius="md" withBorder>
                <Text size="lg">Welcome back!</Text>
            </Card>

            <Card shadow="sm" p="lg" radius="md" withBorder mt="xl">
                <Title order={3}>Your Orders</Title>
                <Table mt="md">
                    <thead>
                    <tr>
                        <th>Product Name</th>
                        <th>Client Name</th>
                        <th>Price</th>
                        <th>Order ID</th>
                    </tr>
                    </thead>
                    <tbody>
                    {orders.map((order: any) => (
                        <tr key={order.id}>
                            <td>{order.productName}</td>
                            <td>{order.clientName}</td>
                            <td>{order.price}</td>
                            <td>{order.id}</td>
                        </tr>
                    ))}
                    </tbody>
                </Table>
                <Pagination
                    page={currentPage}
                    onChange={setCurrentPage}
                    total={Math.ceil(totalOrders / 10)}  // Assuming 10 orders per page
                    mt="md"
                />
            </Card>
        </Container>
    );
}
