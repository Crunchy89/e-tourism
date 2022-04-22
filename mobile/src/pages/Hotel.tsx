import React from 'react';
import {View,Text,StyleSheet,Button} from 'react-native';
import Container from '../component/Container';
import ButtonBack from '../component/ButtonBack';

const Hotel = ({navigation}) => {
    return (
        <>
        <ButtonBack link="Home" navigation={navigation} title="Hotel"/>
        <Container>
            <View>
                <Text>Hotel</Text>
            </View>
        </Container>
        </>
    );
};

export default Hotel;