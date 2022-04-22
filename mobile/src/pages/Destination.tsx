import React from 'react';
import {View,Text,StyleSheet,Button} from 'react-native';
import Container from '../component/Container';
import ButtonBack from '../component/ButtonBack';

const Destination = ({navigation}) => {
    return (
        <>
        <ButtonBack link="Home" navigation={navigation} title="Destination"/>
        <Container>
            <View>
                <Text>Destination</Text>
            </View>
        </Container>
        </>
    );
};

export default Destination;