import React from 'react';
import {StyleSheet,ScrollView, FlatList,View,Image} from 'react-native';
import Container from '../component/Container';
import ButtonBack from '../component/ButtonBack';
import Header from '../component/Header';
import CardImage from "../component/CardImage";

const dummy=[
    {
        id:1,
        image:'https://tempatwisataunik.com/wp-content/uploads/2018/04/56a41e3c-gunung-rinjani.jpg',
        title:'Rinjani',
        rating:4,
    },
    {
        id:2,
        image:'https://tempatwisataunik.com/wp-content/uploads/2018/04/56a41e3c-gunung-rinjani.jpg',
        title:'Gili Terawangan',
        rating:4,
    },
    {
        id:3,
        image:'https://tempatwisataunik.com/wp-content/uploads/2018/04/56a41e3c-gunung-rinjani.jpg',
        title:'Gili Air',
        rating:4,
    },
    {
        id:4,
        image:'https://tempatwisataunik.com/wp-content/uploads/2018/04/56a41e3c-gunung-rinjani.jpg',
        title:'Gili Meno',
        rating:4,
    },
    {
        id:5,
        image:'https://tempatwisataunik.com/wp-content/uploads/2018/04/56a41e3c-gunung-rinjani.jpg',
        title:'Rinjani',
        rating:4,
    },
    {
        id:6,
        image:'https://tempatwisataunik.com/wp-content/uploads/2018/04/56a41e3c-gunung-rinjani.jpg',
        title:'Gili Terawangan',
        rating:4,
    },
    {
        id:7,
        image:'https://tempatwisataunik.com/wp-content/uploads/2018/04/56a41e3c-gunung-rinjani.jpg',
        title:'Gili Air',
        rating:4,
    },
    {
        id:8,
        image:'https://tempatwisataunik.com/wp-content/uploads/2018/04/56a41e3c-gunung-rinjani.jpg',
        title:'Gili Meno',
        rating:4,
    },

];

const Destination = ({navigation}) => {
    return (
        <>
        <ButtonBack link="Home" navigation={navigation} title="Destination"/>
        <Container>
                <ScrollView 
                style={styles.container} 
                showsVerticalScrollIndicator={false}
                showsHorizontalScrollIndicator={false}
                >
                    <View>
                    <Header title="Popular Destination"/>
                    <FlatList
                        horizontal
                        data={dummy}
                        renderItem={({item})=>(
                            <CardImage
                                id={item.id}
                                link={`DetailDestination`}
                                image={item.image}
                                title={item.title}
                                navigation={navigation}
                                />
                                )}
                                keyExtractor={(item) => item.id.toString()}
                                showsHorizontalScrollIndicator={false}
                                />
                    <Header title="New Destination"/>
                    <FlatList
                        horizontal
                        data={dummy}
                        renderItem={({item})=>(
                            <CardImage
                                id={item.id}
                                link={`DetailDestination`}
                                image={item.image}
                                title={item.title}
                                navigation={navigation}
                                />
                                )}
                                keyExtractor={(item) => item.id.toString()}
                                showsHorizontalScrollIndicator={false}
                                />
                    </View>
                    <Header title="Find More"/>
                    <View style={styles.view}>
                    {dummy.map((item,index)=>(
                        <CardImage
                            key={index}
                            id={item.id}
                            link={`DetailDestination`}
                            image={item.image}
                            title={item.title}
                            navigation={navigation}
                        />
                    ))}
                    </View>
                </ScrollView>
        </Container>
        </>
    );
};

const styles = StyleSheet.create({
    container: {
        flex: 1,
        borderRadius:5,
        padding: 10,
        width:"100%",
        backgroundColor: 'rgba(255,255,255,.7)',
        overflow: 'hidden',
    },
    view:{
        flexDirection:'row',
        flexWrap:'wrap',
        justifyContent:'space-around',
        marginBottom:20
    }
});

export default Destination;