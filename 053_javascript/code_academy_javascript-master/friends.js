var friends = {};

friends.bill = {};
friends.steve = {};

friends.bill.firstName = "Bill";
friends.bill.lastName = "Gatttttttttttes";
friends.bill.number = "555-567-7777";
friends.bill.address = ["One Microsoft Way","Redmond","WA","98052"];

friends.steve.firstName = "Steve";
friends.steve.lastName = "Jobbbbbbbbbbbbbs";
friends.steve.number = "444-321-1111";
friends.steve.address = ["1 Infinite Loop","Cupertino","CA","95014"];

//display item in object
// eg, displays friend in friends
var list = function(object) {
    for (var item in object) {
        console.log(item);
    }
};

list(friends);

//find friend

var search = function(name) {
    for(var friend in friends) {
        if(friends[friend].firstName === name) {
            console.log(friends[friend]);
            return friends[friend];
        }
    }
};

search("Bill");