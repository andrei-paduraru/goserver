var controllers = angular.module('controllers', []);
var services = angular.module('services', ['ngResource']);
var app = angular.module('app', [
        'controllers',
        'services'
    ]
);

app.run(
    ['$rootScope', function ($rootScope) {
        console.log("app running");
    }]
);