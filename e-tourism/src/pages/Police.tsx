import React from 'react';
import {View,Text} from 'react-native';
import Container from '../component/Container';
import ButtonBack from '../component/ButtonBack';

const Police = ({navigation}) => {
    return (
        <>
        <ButtonBack link="Home" navigation={navigation} title="Police Station"/>
        <Container>
            <View>
                <Text>Police Station</Text>
            </View>
        </Container>
        </>
    );
};

export default Police;