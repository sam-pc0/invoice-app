import * as React from 'react';
import Button from '@mui/material/Button';
import AppBarComponent from '../../components/AppBar';
import CardProduct from '../../components/CardProduct';
import Grid from '@mui/material/Grid';
import Chip from '@mui/material/Chip';
import { useNavigate } from "react-router-dom";


export default function MyProfile() {
    const navigate = useNavigate();

    const onClick = () => {
        navigate("/");
    };

    const featuredPosts = [
        {
            title: 'Logitech MK345 combo inalámbrico',
            date: 'US$39.99',
            description:
                'Teclado de computadora de tamaño completo con reposamanos. ',
            image: 'https://m.media-amazon.com/images/I/61S0sV1a57L._AC_SL1500_.jpg',
            imageLabel: 'Image Text',
        },
        {
            title: 'Logitech Cámara web C920S HD Pro',
            date: 'US$58.49',
            description:
                'Videollamadas Full HD 1080P: calidad de vídeo premium que te hace parecer un profesional.',
            image: 'https://m.media-amazon.com/images/I/61-6uAf8soL._AC_SL1500_.jpg',
            imageLabel: 'Image Text',
        },



        {
            title: 'SAMSUNG QLED 4K UHD',
            date: 'US$1,897.99',
            description:
                'TV inteligente con múltiples asistentes de voz.',
            image: 'https://m.media-amazon.com/images/I/81inMIbTO9L._AC_SL1500_.jpg',
            imageLabel: 'Image Text',
        },
    ];
    return (
        <>
            <AppBarComponent />
            <Grid container >
                <Grid xs={6} style={{ paddingLeft: '150px', paddingTop: '100px' }} >
                    <Grid xs={3}>
                        <h1>Nombre:</h1>
                    </Grid>
                    <Grid xs={3}>
                        <Chip style={{ fontSize: 18 }} size='medium' label="Alejandro" />
                    </Grid>
                    <Grid xs={3}>
                        <h1>Apellido:</h1>
                    </Grid>
                    <Grid xs={3}>
                        <Chip style={{ fontSize: 18 }} label="Sazo" />
                    </Grid>

                    <Button onClick={onClick} style={{ background: "#2292A4", marginTop: '190px', marginLeft: '190px' }} variant="contained">Cerrar Sesión</Button>
                </Grid>
                <Grid xs={6} style={{ paddingTop: '50px', paddingLeft: '40px' }}>
                    <h1>Ofertas publicadas </h1>
                    <Grid spacing={4}>
                        {featuredPosts.map((post) => (
                            <Grid item xs={12} md={9}>
                                <CardProduct key={post.title} post={post} />
                            </Grid>
                        ))}
                    </Grid>
                </Grid>
            </Grid>
        </>
    );
}