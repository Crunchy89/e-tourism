import React from 'react';
import {View,Text} from 'react-native';
import Container from '../component/Container';
import ButtonBack from '../component/ButtonBack';

const Event = ({navigation}) => {
    return (
        <>
        <ButtonBack link="Home" navigation={navigation} title="Event"/>
        <Container>
            <View>
                <Text>Event</Text>
            </View>
        </Container>
        </>
    );
};

export default Event;