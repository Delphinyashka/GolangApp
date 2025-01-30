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
    const router = useRouter();

    useEffect(() => {
        const storedToken = Cookies.get("jwt");
        if (storedToken) {
            fetchOrders(storedToken, currentPage);
        } else {
            router.push("/sign-in"); // Redirect to login if no token
        }
    }, [currentPage, router]);

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
            setOrders(data.orders);
            setTotalOrders(data.total);
        }
    };

    const handleSignOut = async () => {
        await fetch("http://localhost:8081/user/logout", {
            method: "POST",
            credentials: "include",
        });

        Cookies.remove("jwt"); // Remove the JWT cookie manually
        router.push("/sign-in"); // Redirect to login after logout
    };

    return (
        <Container size="sm" mt={50}>

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
                    total={Math.ceil(totalOrders / 10)}
                    mt="md"
                />
            </Card>

            <Card shadow="sm" p="lg" radius="md" withBorder>
                <Button variant="outline" color="red" size="xl" radius="xl" onClick={handleSignOut}>
                    Logout
                </Button>
            </Card>
        </Container>
    );
}
