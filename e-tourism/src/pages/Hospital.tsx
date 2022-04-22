import React from 'react';
import {View,Text} from 'react-native';
import Container from '../component/Container';
import ButtonBack from '../component/ButtonBack';

const Hospital = ({navigation}) => {
    return (
        <>
        <ButtonBack link="Home" navigation={navigation} title="Hospital"/>
        <Container>
            <View>
                <Text>Hospital</Text>
            </View>
        </Container>
        </>
    );
};

export default Hospital;