// a fuzzer for undocumented params

const AWS = require('aws-sdk');
const fetch = require('node-fetch');

var credentials = new AWS.SharedIniFileCredentials({profile: 'nothing'});
AWS.config.credentials = credentials;
AWS.config.region = 'us-east-1';

var long_undocumented_test_list = `
APIGateway.CreateAuthorizer {3} - apigateway
APIGateway.CreateDocumentationPart {3} - apigateway
APIGateway.ImportRestApi {3} - apigateway
APIGateway.PutRestApi {3} - apigateway
APIGateway.TagResource {3} - apigateway
APIGateway.UpdateRestApi {3} - apigateway
CognitoIdentityServiceProvider.RevokeToken {3} - cognitoidp
DirectConnect.ConfirmCustomerAgreement {3} - directconnect
DirectConnect.DescribeCustomerMetadata {3} - directconnect
DirectConnect.DescribeRouterConfiguration {3} - directconnect
DirectConnect.UpdateDirectConnectGateway {3} - directconnect
DirectoryService.DescribeClientAuthenticationSettings {3} - ds
DynamoDB.DescribeEndpoints {3} - dynamodb
EC2.DescribeCapacityReservationFleets {3} - ec2
EC2.GetInstanceTypesFromInstanceRequirements {3} - ec2
EC2.GetSpotPlacementScores {3} - ec2
Rekognition.CreateDataset {3} - rekognition
Rekognition.DeleteDataset {3} - rekognition
Rekognition.DescribeDataset {3} - rekognition
Rekognition.DistributeDatasetEntries {3} - rekognition
Rekognition.ListDatasetEntries {3} - rekognition
Rekognition.ListDatasetLabels {3} - rekognition
Rekognition.UpdateDatasetEntries {3} - rekognition
StorageGateway.UpdateSMBLocalGroups {3} - storagegateway
MediaLive.ClaimDevice {3} - medialive
SageMaker.BatchDescribeModelPackage {3} - sagemaker
SageMaker.CreateStudioLifecycleConfig {3} - sagemaker
SageMaker.RetryPipelineExecution {3} - sagemaker
SageMaker.UpdateProject {3} - sagemaker
WorkMail.DeleteMobileDeviceAccessOverride {3} - workmail
WorkMail.DeregisterMailDomain {3} - workmail
WorkMail.GetMailDomain {3} - workmail
WorkMail.GetMobileDeviceAccessOverride {3} - workmail
WorkMail.ListMailDomains {3} - workmail
WorkMail.ListMobileDeviceAccessOverrides {3} - workmail
WorkMail.PutMobileDeviceAccessOverride {3} - workmail
WorkMail.RegisterMailDomain {3} - workmail
WorkMail.UpdateDefaultMailDomain {3} - workmail
TranscribeService.ListTagsForResource {3} - transcribe
TranscribeService.TagResource {3} - transcribe
TranscribeService.UntagResource {3} - transcribe
Connect.StartContactStreaming {3} - connect
Connect.StopContactStreaming {3} - connect
MediaTailor.ConfigureLogsForPlaybackConfiguration {3} - mediatailor
MediaTailor.CreatePrefetchSchedule {3} - mediatailor
MediaTailor.DeletePrefetchSchedule {3} - mediatailor
MediaTailor.GetPrefetchSchedule {3} - mediatailor
MediaTailor.ListPrefetchSchedules {3} - mediatailor
EKS.RegisterCluster {3} - eks
Route53Resolver.GetResolverConfig {3} - route53resolver
Route53Resolver.ListResolverConfigs {3} - route53resolver
Route53Resolver.UpdateResolverConfig {3} - route53resolver
SecurityHub.CreateFindingAggregator {3} - securityhub
SecurityHub.DeleteFindingAggregator {3} - securityhub
SecurityHub.GetFindingAggregator {3} - securityhub
SecurityHub.ListFindingAggregators {3} - securityhub
SecurityHub.UpdateFindingAggregator {3} - securityhub
ApiGatewayV2.CreateAuthorizer {3} - apigateway
ApiGatewayV2.CreateVpcLink {3} - apigateway
WAFV2.GetManagedRuleSet {3} - wafv2
WAFV2.ListAvailableManagedRuleGroupVersions {3} - wafv2
WAFV2.ListManagedRuleSets {3} - wafv2
WAFV2.PutManagedRuleSetVersions {3} - wafv2
WAFV2.UpdateManagedRuleSetVersionExpiryDate {3} - wafv2
NetworkManager.GetNetworkResourceCounts {3} - networkmanager
NetworkManager.GetNetworkResourceRelationships {3} - networkmanager
NetworkManager.GetNetworkResources {3} - networkmanager
NetworkManager.GetNetworkRoutes {3} - networkmanager
NetworkManager.GetNetworkTelemetry {3} - networkmanager
NetworkManager.GetRouteAnalysis {3} - networkmanager
NetworkManager.StartRouteAnalysis {3} - networkmanager
NetworkManager.UpdateNetworkResourceMetadata {3} - networkmanager
Outposts.CreateOrder {3} - outposts
LexModelsV2.ListAggregatedUtterances {3} - lex
SSMContacts.GetContactPolicy {1}
SSMContacts.ListTagsForResource {1}
`;

var found_permissions = [];

function transformArn(arn) {
    return arn
        .replace(/PN/g, "$\{")
        .replace(/XX/g, "}")
        .replace(/pn/g, "$\{")
        .replace(/xx/g, "}")
        .replace(/774857101424/g, "${Account}")
        .replace(/us-east-1/g, "${Region}")
        .replace(/arn:aws/g, "arn:${Partition}");
}

async function go() {
    var known_permissions = {};
    var iamdefdata = await fetch('https://raw.githubusercontent.com/iann0036/iam-dataset/main/iam_definition.json');
    iamdef = await iamdefdata.json();

    for (var iamdefitem of iamdef) {
        for (var priv of iamdefitem.privileges) {
            known_permissions[iamdefitem.prefix+":"+priv.privilege] = (priv.access_level == "Unknown");
        }
    }
    
    for (let test_item of long_undocumented_test_list.split("\n")) {
        test_item = test_item.split(" ")[0];
        if (test_item != "") {
            let split_test_item = test_item.split(".");
            let svc = new AWS[split_test_item[0]]({});
            let method = split_test_item[1][0].toLowerCase() + split_test_item[1].substr(1);

            let complete = false;
            let params = {};
            let iterations = 0;
            let last_param = null;

            while (!complete) {
                try {
                    console.log("--\n" + method);
                    console.log(params);
                    await svc[method].call(svc, params).promise();
                    complete = true;
                } catch (err) {
                    if (err.message) {
                        if (err.message.includes("not authorized to perform: ")) {
                            console.log(err.message);
                            let match = err.message.match(/not authorized to perform: ([a-zA-Z0-9-:]+)(?: on resource: (.+))?/);
                            let permission = match[1];
                            let resource = match[2];

                            found_permissions.push({
                                'permission': permission,
                                'resource': resource || null,
                                'service': split_test_item[0],
                                'method': split_test_item[1]
                            });
                            complete = true;
                        } else {
                            if (err.message.includes("Missing required key '")) {
                                let paramname = err.message.match(/Missing required key '(.+)'/)[1];
                                params[paramname] = "PN"+paramname+"XX";
                                if (paramname.toLowerCase().includes("accountid")) {
                                    params[paramname] = "774857101424";
                                }
                                last_param = paramname;
                            } else if (err.message.includes(" to be an Array")) {
                                params[last_param] = [];
                            } else if (err.message.includes(" to be a structure")) {
                                params[last_param] = {};
                            } else if (err.message.includes(" to be a number")) {
                                params[last_param] = 1;
                            } else if (err.message.includes(" to be a boolean")) {
                                params[last_param] = false;
                            } else {
                                console.log(err.message);
                                complete = true;
                            }
                        }
                    } else {
                        complete = true;
                    }
                }
                iterations += 1;
                if (iterations > 20) {
                    complete = true;
                }
            }
        }
    }

    console.log(found_permissions);

    let res = {};
    for (let item of found_permissions) {
        if (!Object.keys(known_permissions).includes(item['permission'])) {
            res[item['service'] + "." + item['method']] = [{
                "action": item['permission'],
                "undocumented": true
            }];
            if (item['resource']) {
                res[item['service'] + "." + item['method']][0]['arn_override'] = {
                    "template": transformArn(item['resource'])
                };
            }
        } else {
            if (!known_permissions[item['permission']]) {
                console.log("Invalid hit: " + item['permission']);
            } else {
                res[item['service'] + "." + item['method']] = [{
                    "action": item['permission'],
                    "undocumented": true
                }];
                if (item['resource']) {
                    res[item['service'] + "." + item['method']][0]['arn_override'] = {
                        "template": transformArn(item['resource'])
                    };
                }
            }
        }
    }

    var mapdata = await fetch('https://raw.githubusercontent.com/iann0036/iam-dataset/main/map.json');
    var map = await mapdata.json();

    for (let reskey of Object.keys(res)) {
        if (Object.keys(map['sdk_method_iam_mappings']).includes(reskey)) {
            delete res[reskey];
        }
    }

    console.log(JSON.stringify(res, null, 4));
}

go();
