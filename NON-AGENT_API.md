# NON-AGENT API DOCUMENT

**Started:** 2008-07-24  
**Updated:** 2025-07-18

This document describes the functions of an API (Application Programming Interface) for all functions NOT directly relating to the VICIDIAL Agent screen. This functionality will be rather limited at first and will be built upon as critical functions are identified and programmed into it.

There is also an Agent API script, for more information on that, please read the `AGENT_API.txt` document.

> [!NOTE]
> The API is able to use HTTPS if your webserver is configured to use HTTPS. For example, all VICIhost accounts are able to use HTTPS for API functions.

---

## API Functions

*   **version** - shows version and build of the API, along with the date/time
*   **sounds_list** - outputs a list of audio files from the audio store
*   **moh_list** - outputs a list of music on hold classes in the system
*   **vm_list** - outputs a list of voicemail boxes in the system
*   **blind_monitor** - calls user-defined phone and places them in session as blind monitor
*   **add_lead** - adds a new lead to the vicidial_list table with several fields and options
*   **update_lead** - updates lead information in vicidial_list and associated custom table
*   **batch_update_lead** - updates multiple leads in vicidial_list with single set of field values (limited field options)
*   **agent_ingroup_info** - shows in-group and outbound auto-dial info for logged-in agent
*   **recording_lookup** - looks up recordings based upon user and date or lead_id
*   **did_log_export** - exports all calls inbound to a DID for one day
*   **agent_stats_export** - exports agent statistics for set time period
*   **add_user** - adds a user to the system
*   **update_user** - updates some user settings in the system
*   **copy_user** - copies an existing user in the system to a new user ID and name
*   **add_phone** - adds a phone to the system
*   **update_phone** - updates or deletes an existing phone record in the system
*   **add_phone_alias** - adds a phone alias record to the system
*   **update_phone_alias** - updates or deletes an existing phone alias record in the system
*   **add_list** - adds a list to the system
*   **update_list** - updates list settings in the system, reset leads in list, delete list
*   **list_info** - summary information about a list
*   **list_custom_fields** - shows the custom fields that are in a list, or all lists
*   **add_group_alias** - adds group alias to the system
*   **user_group_status** - real-time status of one or more user groups
*   **in_group_status** - real-time status of one or more in groups
*   **agent_status** - real-time status of one user
*   **callid_info** - information about a call based upon the caller_code or call ID
*   **lead_field_info** - pulls the value of one field of a lead
*   **lead_callback_info** - outputs scheduled callback data for a specific lead
*   **lead_all_info** - outputs all lead data for a lead(optionally including custom fields)
*   **lead_status_search** - displays all field values of all leads that match the status and call date in request
*   **lead_search** - searches for leads in the vicidial_list table by phone numbers
*   **ccc_lead_info** - outputs lead data for cross-cluster-communication call
*   **server_refresh** - forces a conf file refresh on all telco servers in the cluster
*   **check_phone_number** - allows you to check if a phone number is valid and dialable
*   **logged_in_agents** - list of agents that are logged in to the system
*   **update_campaign** - updates campaign settings in the system
*   **add_did** - adds new Inbound DID entries to the system
*   **copy_did** - adds new Inbound DID entries to the system, copied from an existing DID entry
*   **update_did** - updates Inbound DID information in the system
*   **phone_number_log** - exports list of calls placed to one of more phone numbers
*   **call_status_stats** - report on number of calls made by campaign and ingroup, with hourly and status breakdowns
*   **call_dispo_report** - call disposition breakdown report
*   **agent_campaigns** - looks up allowed campaigns/in-groups for a specific user
*   **update_cid_group_entry** - updates CID Group entries
*   **add_dnc_phone** - adds a phone number to one of the DNC lists
*   **delete_dnc_phone** - removes a phone number from one of the DNC lists
*   **add_fpg_phone** - adds a phone number to a Filter Phone Group
*   **delete_fpg_phone** - removes a phone number from a Filter Phone Group
*   **campaigns_list** - displays information about all campaigns in the system
*   **hopper_list** - displays information about leads in the hopper for a campaign
*   **hopper_bulk_insert** - adds group of lead_ids to the vicidial_hopper table
*   **update_alt_url** - updates/adds/displays alternate dispo call url entries for a campaign
*   **update_presets** - updates/adds/displays/deletes campaign preset entries
*   **lead_dearchive** - moves a lead from the archive to active leads table
*   **update_remote_agent** - updates some remote agent settings in the system
*   **user_details** - display information about one user

---

## New Scripts

*   `/vicidial/non_agent_api.php` - the script that is accessed to execute commands

---

## Changelog

| Date/Build | Description |
| :--- | :--- |
| **80724-1723** | First build |
| **90428-0208** | Added `blind_monitor` function |
| **90513-1720** | Added `sounds_list` function |
| **90721-1432** | Added rank and owner |
| **90904-1535** | Added `moh_list` musiconhold function |
| **90916-2342** | Added `vm_list` voicemail list |
| **91203-1140** | Added `agent_ingroup_info` feature |
| **91216-0331** | Added duplication check features to `add_lead` function |
| **100118-0543** | Added new Australian and New Zealand DST schemes (FSO-FSA and LSS-FSA) |
| **100704-1148** | Added custom fields inserts to the `add_lead` function |
| **100718-0245** | Added `update_lead` function to update existing leads |
| **100723-1333** | Added `no_update` option to the `update_lead` function |
| **100728-1952** | Added `delete_lead` option to the `update_lead` function |
| **100924-1403** | Added `called_count` as an `update_lead` option |
| **101111-1536** | Added `vicidial_hopper.source` to `vicidial_hopper` inserts |
| **101117-1104** | Added callback and custom field entry delete to `delete_lead` option |
| **101206-2126** | Added `recording_lookup` and `did_log_export` functions |
| **110127-2245** | Added `add_user` and `add_phone` functions |
| **110306-1044** | Added `add_list` and `update_list` functions |
| **110316-2035** | Added `reset_time` variable and `NAMEPHONE` dup search |
| **110404-1356** | Added `uniqueid` search parameter to `recording_lookup` function |
| **110409-0822** | Added run_time logging of API functions |
| **110424-0854** | Added option for time zone code lookups using owner field in `add_lead` function |
| **110529-1220** | Added time zone information output to `version` function |
| **110614-0726** | Added `reset_lead` option to `update_lead` function (issue #502) |
| **110705-1928** | Added options for USACAN 4th digit prefix (no 0 or 1) and valid areacode filtering to `add_lead` |
| **110821-2318** | Added `update_phone`, `add_phone_alias`, `update_phone_alias` functions |
| **110928-2110** | Added callback options to `add_lead` and `update_lead` |
| **111106-0959** | Added `user_group` restrictions to some functions |
| **120127-1331** | Small fix for plus replacement in custom fields strings for `add`/`update_lead` functions |
| **120210-1215** | Small change for hopper adding `vendor_lead_code` |
| **120213-1613** | Added optional logging of all non-admin.php requests, enabled in `options.php` |
| **120326-1307** | Added `agent_stats_export` function |
| **120810-0859** | Added `add_group_alias` function, altered `agent_stats_export` function |
| **120912-2042** | Added `user_group_status` and `in_group_status` functions |
| **120913-1255** | Added `update_log_entry` function |
| **121116-1938** | Added state call time restrictions to `add_lead` hopper insert function |
| **121125-2210** | Added Other Campaign DNC option and list expiration date option |
| **130328-0949** | Added `update_phone_number` option to `update_lead` function |
| **130405-1538** | Added `agent_status` function |
| **130420-1938** | Added NANPA prefix validation and timezone options |
| **130614-0907** | Added pause code to output of `agent_status` function |
| **130617-2232** | Added real-time sub-statuses to output of `agent_status` function |
| **140124-1057** | Added `callid_info` function |
| **140206-1205** | Added `update_user` function |
| **140211-1056** | Added `server_refresh` function |
| **140214-1540** | Added `check_phone_number` function |
| **140331-2119** | Converted division calculations to use MathZDC function |
| **140403-2024** | Added `camp_rg_only` option to `update_user` function |
| **140418-1553** | Added `preview_lead_id` for `agent_status` |
| **140617-2029** | Added `vicidial_users` `wrapup_seconds_override` option |
| **140812-0939** | Added `phone_number` and `vendor_lead_code` to `agent_status` function output |
| **150123-1611** | Fixed issue with list local call times and `add_lead` |
| **150217-1404** | Archive deleted callbacks to `vicidial_callbacks_archive` table |
| **150309-0250** | Added ability to use urlencoded web form addresses |
| **150313-0818** | Allow for single quotes in `vicidial_list` and custom data fields |
| **150428-1720** | Added `web_form_address_three` to `add_list`/`update_list` functions |
| **150430-0644** | Added API allowed function restrictions and allowed list restrictions |
| **150512-2028** | Added filtering of hash sign on some input variables, Issue #851 |
| **150603-1528** | Fixed format issue in `recording_lookup` |
| **150730-2022** | Added option to set `entry_list_id` |
| **150804-0948** | Added `WHISPER` option for `blind_monitor` function |
| **150808-1438** | Added compatibility for custom fields data option |
| **151226-0954** | Added `session_id(conf_exten)` field to output of `agent_status`, and added `logged_in_agents` function |
| **160603-1041** | Added `update_campaign` function |
| **160709-2219** | Added `lead_field_info` function |
| **161020-1042** | Added `lookup_state` option to `add_lead`, added 10-digit validation if `usacan_areacode_check` enabled |
| **170204-0947** | Added `phone_number_log` function |
| **170209-1149** | Added URL and IP logging |
| **170219-1520** | Added 90-day duplicate check option |
| **170223-0743** | Added QXZ translation to admin web user functions text |
| **170301-1649** | Added validation that sounds web dir exists to `sounds_list` function |
| **170409-1603** | Added IP List validation code |
| **170423-0800** | Added `force_entry_list_id` option for `update_lead` |
| **170508-1048** | Added `blind_monitor` logging |
| **170527-2254** | Fix for rare inbound logging issue #1017 |
| **170601-0747** | Added `add_to_hopper` options to `update_lead` function |
| **170609-1107** | Added `ccc_lead_info` function |
| **170615-0006** | Added `DIAL` status for manual dial agent calls that have not been answered, to 4 functions |
| **171229-1602** | Added `lead_status_search` function |
| **180109-1313** | Added `call_status_stats` function |
| **180208-1655** | Added `call_dispo_report` function |
| **180331-1716** | Added options to `update_campaign`, `add_user`, `update_phone`. Added function `update_did` |
| **180916-1059** | Added check for daily list reset limit |
| **181214-1606** | Added `lead_callback_info` function |
| **190325-1055** | Added `agent_campaigns` function |
| **190628-1504** | Added `update_cid_group_entry` function |
| **190703-0858** | Added `custom_fields_copy` option to `add_list` function |
| **190930-1629** | Added `list_description` field to `add_list` and `update_list` functions |
| **191107-1143** | Added `list_info` function |
| **191113-1758** | Added `add_dnc_phone` and `add_fpg_phone` functions |
| **191114-1637** | Added `remove_from_hopper` option to `update_lead` function |
| **200331-1207** | Added `list_custom_fields` function |
| **200418-0837** | Added `custom_copy_method` option for the `add_list` function |
| **200606-0938** | Added more settings to `add_list` & `update_list` |
| **200610-1447** | Added duration option to `recording_lookup` function |
| **200622-1609** | Added more options to `add_phone` and `update_phone` functions |
| **200815-1025** | Added `campaigns_list` & `hopper_list` functions |
| **200824-2330** | Added search_method `BLOCK` option for `hopper_list` function |
| **201002-1545** | Added extension as `recording_lookup` option, Allowed for secure `sounds_web_server` setting |
| **201106-1654** | Added `campaign_id` option to `agent_stats_export` function |
| **201113-0713** | Added `delete_did` option to `update_did` function |
| **201201-1625** | Added `group_by_campaign` option to `agent_stats_export` function |
| **210113-1714** | Added `copy_user` function |
| **210209-2014** | Added `add_did` function |
| **210210-1627** | Added duplicate check with more X-day options for `add_lead` function |
| **210216-1415** | Added `list_exists_check` option for `add_lead` function |
| **210217-1454** | Added `menu_id` option for `add_did` and `update_did` functions |
| **210227-1116** | Added xferconf number options to `update_campaign`, `add_list` & `update_list` functions |
| **210303-1802** | Added `list_exists_check` option for `update_lead` function |
| **210316-2207** | Added `lead_all_info` function |
| **210319-1720** | Added special 'xDAYS' value option for `callback_datetime` |
| **210320-2127** | Added 'custom_fields_add' option for `update_list` function |
| **210322-1218** | Added display of bad wav file formats on `sounds_list` function |
| **210325-2042** | Added more data output fields to `agent_stats_export` function |
| **210330-1633** | Added ability to use `custom_fields_copy` on `update_list` on a deleted `list_id` |
| **210401-1058** | Added 'custom_fields_update' option for `update_list` function |
| **210402-1102** | Added 'custom_fields_delete' option for `update_list` function |
| **210406-1047** | Added 'dialable_count' option to `list_info` function |
| **210517-1850** | Added `phone_code` as modifiable field in `update_lead` function |
| **210611-1610** | Added more variables for `add_did`/`update_did` functions |
| **210618-1000** | Added CORS support |
| **210625-1421** | Added `BARGESWAP` option to `blind_monitor` function |
| **210701-2050** | Added `ingroup_rank`/`ingroup_grade` options to the `update_user` function |
| **210812-2640** | Added update of live agent records for `ingroup_rank`/`ingroup_grade` in the `update_user` function |
| **210917-1615** | Added `batch_update_lead` function |
| **211107-1556** | Added optional `phone_number` search for `lead_all_info` function |
| **211118-1946** | Added 'delete_cf_data' setting to the `update_lead` function |
| **220126-2116** | Added `dispo_call_url` and alt URL options for `update_campaign`. Added `update_alt_url` function |
| **220307-0937** | Added `update_presets` function |
| **220310-0825** | Added DELETE action to the `update_presets` function |
| **220902-0823** | Added `dial_status_add`/`dial_status_remove` options to `update_campaign` function |
| **220920-0814** | Added more "XDAY" options to `add_lead` duplicate checks |
| **221108-1849** | Added `include_ip` option for `agent_status` function |
| **230122-1821** | Added `reset_password` option to `update_user` function |
| **230308-1758** | Fix for `update_lead` list-restrict phone number update issue |
| **230614-0828** | Added `container_list` function |
| **230721-1753** | Added insert for daily RT monitoring |
| **231109-0933** | Added `archived_lead` option to multiple functions |
| **231109-2022** | Added `lead_dearchive` function |
| **240101-0920** | Added `DUPPHONEALT...` options for duplicate_check within `add_lead` function |
| **240105-1009** | Added `delete_fpg_phone` function |
| **240120-0840** | Added `list_order`, `list_order_randomize`, `list_order_secondary` optional fields to `update_campaign` function |
| **240217-0907** | Added more missing `vicidial_users` fields from copy function |
| **240302-0804** | Added `delete_dnc_phone` function |
| **240703-1657** | Added `update_remote_agent` function |
| **240718-0942** | Fixes for inconsistencies in documentation and permissions for some functions |
| **240718-1716** | Added fields to `campaigns_list` function output |
| **240730-1832** | Changes for PHP8 compatibility, Added `copy_did` function |
| **240824-1626** | Added `user_details` function |
| **241004-1518** | Added `webform_one`-`three` variables to the `update_campaign` function |
| **241113-1600** | Added `in_groups` as input option for `update_user` function |
| **250720-1841** | Added `hopper_bulk_insert` function |

---

### General Notes
API Functions use the `function` variable.

Just as with the Agent API, the non-agent API requires the user and pass of a valid api-enabled `vicidial_users` account to execute actions.

#### Example Auth Error Response
```text
ERROR: auth USER DOES NOT HAVE PERMISSION TO USE THIS FUNCTION - 6666|add_lead
```

---

## Functions

### version
Shows version and build of the API, along with the date/time and timezone.

#### Example URL
```http
http://server/vicidial/non_agent_api.php?function=version
```

#### Example Response
```text
VERSION: 2.4-34|BUILD: 110424-0854|DATE: 2011-05-29 12:19:22|EPOCH: 1306685962|DST: 1|TZ: -5|TZNOW: -4| 
```

---

### sounds_list
Outputs a list of audio files from the audio store.

> [!NOTE]
> API user for this function must have `user_level` set to 7 or higher.

#### Optional Fields
*   `format` - format of the output (`tab`, `link`, `selectframe`)
*   `stage` - how to sort the output (`date`, `size`, `name`)
*   `comments` - name of the field to populate

#### Example URL
```http
http://server/vicidial/non_agent_api.php?source=test&user=6666&pass=1234&function=sounds_list&format=selectframe&comments=fieldname&stage=date
```

#### Example Responses
*(Success is inferred by output of audio files list)*

```text
ERROR: sounds_list USER DOES NOT HAVE PERMISSION TO VIEW SOUNDS LIST
ERROR: audio store web directory does not exist: |6666|sounds_list|/srv/www/htdocs/1234|
ERROR: sounds_list CENTRAL SOUND CONTROL IS NOT ACTIVE
```

---

### moh_list
Outputs a list of music on hold classes in the system.

> [!NOTE]
> API user for this function must have `user_level` set to 7 or higher.

#### Optional Fields
*   `format` - format of the output (only `selectframe` works)
*   `comments` - name of the field to populate

#### Example URL
```http
http://server/vicidial/non_agent_api.php?source=test&user=6666&pass=1234&function=moh_list&format=selectframe&comments=fieldname&stage=date
```

#### Example Responses
*(Success is inferred by output of MOH classes list)*

```text
ERROR: moh_list USER DOES NOT HAVE PERMISSION TO VIEW SOUNDS LIST
ERROR: moh_list CENTRAL SOUND CONTROL IS NOT ACTIVE
```

---

### vm_list
Outputs a list of voicemail boxes in the system.

> [!NOTE]
> API user for this function must have `user_level` set to 7 or higher.

#### Optional Fields
*   `format` - format of the output (only `selectframe` works)
*   `comments` - name of the field to populate

#### Example URL
```http
http://server/vicidial/non_agent_api.php?source=test&user=6666&pass=1234&function=vm_list&format=selectframe&comments=fieldname&stage=date
```

#### Example Responses
*(Success is inferred by output of voicemail boxes list)*

```text
ERROR: vm_list USER DOES NOT HAVE PERMISSION TO VIEW VOICEMAIL BOXES LIST
ERROR: vm_list CENTRAL SOUND CONTROL IS NOT ACTIVE
```

---

### blind_monitor
Calls user-defined phone and places them in session as blind monitor.

> [!NOTE]
> API user for this function must have `user_level` set to 7 or higher.

#### Required Fields
*   `phone_login` - alpha-numeric, no spaces or special characters allowed
*   `session_id` - must be all numbers, 7 digits
*   `server_ip` - must be all numbers and dots, max 15 characters
*   `source` - description of what originated the API call (maximum 20 characters)
*   `stage` - `MONITOR`, `BARGE` (or `BARGESWAP`) or `HIJACK`, default is `MONITOR` *(HIJACK option is not currently functional)*

#### Example URL
```http
http://server/vicidial/non_agent_api.php?source=test&user=6666&pass=1234&function=blind_monitor&phone_login=350a&session_id=8600051&server_ip=10.10.10.16&stage=MONITOR
```

#### Example Responses
```text
SUCCESS: blind_monitor HAS BEEN LAUNCHED - 350a|010*010*010*017*350|8600051
ERROR: blind_monitor INVALID PHONE LOGIN - 350q 
ERROR: blind_monitor INVALID SESSION ID - 8602051 
ERROR: blind_monitor USER DOES NOT HAVE PERMISSION TO BLIND MONITOR - 6666|0 
ERROR: NO FUNCTION SPECIFIED
```

---

### agent_ingroup_info
Shows in-group and outbound auto-dial info for logged-in agent.

> [!NOTE]
> API user for this function must have `user_level` set to 7 or higher.

#### Required Fields
*   `agent_user` - 2-20 characters
*   `source` - description of what originated the API call (maximum 20 characters)

#### Settings Fields
*   `stage` - `info` (show information only), `change` (show options to change), `text` (standard non-HTML output)

#### Example URL
```http
http://server/vicidial/non_agent_api.php?source=test&function=agent_ingroup_info&stage=change&user=6666&pass=1234&agent_user=1000
```

#### Example Responses
```text
ERROR: agent_ingroup_info USER DOES NOT HAVE PERMISSION TO GET AGENT INFO - 6666|0
ERROR: agent_ingroup_info INVALID USER ID - 1255|6666
```

---

### agent_campaigns
Looks up allowed campaigns/in-groups for a specific user.

> [!NOTE]
> API user for this function must have `user_level` set to 8 or higher, and "view reports" and "modify users" enabled.

#### Required Fields
*   `agent_user` - 2-20 characters
*   `source` - description of what originated the API call (maximum 20 characters)

#### Optional Fields
*   `campaign_id` - 2-8 characters
*   `ignore_agentdirect` - `Y` or `N` (default `N`). Exclude AGENTDIRECT in-groups from results or not.

#### Settings Fields
*   `stage` - the format of the exported data: `csv`, `tab`, `pipe` (default)
*   `header` - include a header (`YES`) or not (`NO`). Optional, default is `NO`.

#### Example URL
```http
http://server/vicidial/non_agent_api.php?source=test&function=agent_campaigns&user=6666&pass=1234&agent_user=1000
http://server/vicidial/non_agent_api.php?source=test&function=agent_campaigns&campaign_id=TESTCAMP&ignore_agentdirect=Y&user=6666&pass=1234&agent_user=1000&stage=csv
```

#### Example Responses
```text
ERROR: agent_campaigns USER DOES NOT HAVE PERMISSION TO GET AGENT INFO - 6666|0
ERROR: agent_campaigns INVALID USER ID - 1255|6666
ERROR: agent_campaigns AGENT USER DOES NOT EXIST - 1000|6666
ERROR: agent_campaigns THIS AGENT USER HAS NO AVAILABLE CAMPAIGNS - 1000|6666
```

##### Successful Output Format
```text
user|allowed_campaigns_list|allowed_ingroups_list
1000|TESTCAMP-TEST123-TEST987|TEST_IN-TEST_IN2-TEST_IN3-TEST_IN4-SALESLINE
```

---

### campaigns_list
Displays information about all campaigns in the system.

> [!NOTE]
> API user for this function must have `user_level` set to 8 or higher and "view reports" enabled.

#### Required Fields
*   `source` - description of what originated the API call (maximum 20 characters)

#### Optional Fields
*   `campaign_id` - 2-8 characters, for all campaigns leave blank

#### Settings Fields
*   `stage` - the format of the exported data: `csv`, `tab`, `pipe` (default)
*   `header` - include a header (`YES`) or not (`NO`). Optional, default is `NO`.

#### Example URL
```http
http://server/vicidial/non_agent_api.php?source=test&function=campaigns_list&user=6666&pass=1234
http://server/vicidial/non_agent_api.php?source=test&function=campaigns_list&campaign_id=TESTCAMP&user=6666&pass=1234&stage=csv
```

#### Example Responses
```text
ERROR: campaigns_list USER DOES NOT HAVE PERMISSION TO GET CAMPAIGN INFO - 6666|0
ERROR: campaigns_list THIS USER HAS NO VIEWABLE CAMPAIGNS - 6666
```

##### Successful Output Format (Header Enabled)
| campaign_id | campaign_name | active | user_group | dial_method | dial_level | lead_order | dial_statuses | dial_timeout | dial_prefix | manual_dial_prefix | three_way_dial_prefix |
| :--- | :--- | :--- | :--- | :--- | :--- | :--- | :--- | :--- | :--- | :--- | :--- |
| TESTCAMP | Test Campaign | Y | ---ALL--- | INBOUND_MAN | 1 | DOWN | PDROP DROP B NEW | 65 | 9 | 9 | 9 |

---

### hopper_list
Displays information about leads in the hopper for a campaign.

> [!NOTE]
> API user for this function must have `user_level` set to 8 or higher and "view reports" enabled.

#### Required Fields
*   `source` - description of what originated the API call (maximum 20 characters)
*   `campaign_id` - 2-8 characters

#### Settings Fields
*   `stage` - the format of the exported data: `csv`, `tab`, `pipe` (default)
*   `header` - include a header (`YES`) or not (`NO`). Optional, default is `NO`.

#### Optional Fields
*   `search_method` - for faster but blocking SQL queries that may affect database performance, set this to `BLOCK`.

#### Example URL
```http
http://server/vicidial/non_agent_api.php?source=test&function=hopper_list&campaign_id=TESTCAMP&user=6666&pass=1234&stage=csv
```

#### Example Responses
```text
ERROR: hopper_list USER DOES NOT HAVE PERMISSION TO GET CAMPAIGN INFO - 6666|0
ERROR: hopper_list THIS CAMPAIGN DOES NOT EXIST - 6666|TESTCAMP
ERROR: hopper_list THERE ARE NO LEADS IN THE HOPPER FOR THIS CAMPAIGN - 6666|TESTCAMP
```

##### Successful Output Format (CSV)
| hopper_order | priority | lead_id | list_id | phone_number | phone_code | state | status | count | gmt_offset | rank | alt | hopper_source | vendor_lead_code | source_id | age_days | last_call_time |
| :--- | :--- | :--- | :--- | :--- | :--- | :--- | :--- | :--- | :--- | :--- | :--- | :--- | :--- | :--- | :--- | :--- |
| 0 | 0 | 796292 | 107 | 9999001095 | 1 | NY | PDROP | 7 | -4.00 | 0 | | S | 1001095 | A321 | 2681 | 1 YEARS |
| 1 | 0 | 796295 | 107 | 9999001098 | 1 | NY | PDROP | 2 | -4.00 | 0 | | S | 1001098 | A322 | 2681 | 3 YEARS |

---

### hopper_bulk_insert
Adds a group of `lead_ids` to the `vicidial_hopper` table.

> [!NOTE]
> API user for this function must have `modify_campaigns` set to 1 and `user_level` must be set to 8 or higher.
> This function will insert leads into the hopper of the campaign that is tied to the `list_id` that each lead belongs to.

#### Required Fields
*   `lead_ids` - must be all numbers and commas (1-9 digit numbers separated by a comma, limit of 1000 `lead_ids`)
*   `source` - description of what originated the API call (maximum 20 characters)

#### Settings Fields
*   `hopper_priority` - `99` to `-99` (higher is prioritized, default is `0`)
*   `hopper_local_call_time_check` - `Y` or `N` (default `N`). Validate the local call time and/or state call time before inserting lead in the hopper.

#### Example URL
```http
http://server/vicidial/non_agent_api.php?source=test&user=6666&pass=1234&function=hopper_bulk_insert&lead_ids=193715,1677922&hopper_priority=97&hopper_local_call_time_check=Y
```

#### Example Responses
```text
NOTICE: hopper_bulk_insert LEAD ADDED TO HOPPER - 7275551111|193715|1677922|6666
NOTICE: hopper_bulk_insert NOT ADDED TO HOPPER, OUTSIDE OF LOCAL TIME - 7275551111|193715|-5|NY|0
NOTICE: hopper_bulk_insert NOT ADDED TO HOPPER, CAMPAIGN DOES NOT EXIST - 101|193715
NOTICE: hopper_bulk_insert NOT ADDED TO HOPPER, LEAD NOT FOUND - 193715
NOTICE: hopper_bulk_insert NOT ADDED TO HOPPER, LEAD IS ALREADY IN THE HOPPER - 193715
NOTICE: hopper_bulk_insert NOT ADDED TO HOPPER, LEAD HOPPER INSERT ERROR - 193715

SUCCESS: hopper_bulk_insert LEADS HAVE BEEN INSERTED INTO THE HOPPER - 2|0|8432948|6666

ERROR: hopper_bulk_insert NO LEADS HAVE BEEN INSERTED INTO THE HOPPER - 1|0|6666
ERROR: hopper_bulk_insert USER DOES NOT HAVE PERMISSION TO ADD LEADS TO THE HOPPER - 6666|0 
```

---

### recording_lookup
Looks up recordings based upon user and date or `lead_id`.

> [!NOTE]
> API user for this function must have `user_level` set to 7 or higher and "view reports" enabled.
> There is a hard limit of 100,000 results.

#### Search Fields
*   `agent_user` - 2-20 characters
*   `lead_id` - 1-10 digits
*   `date` - date of the calls to pull (YYYY-MM-DD format)
*   `uniqueid` - uniqueid of the call (works best included with another search field)
*   `extension` - 3-100 characters, the extension listed in the recording_log

#### Settings Fields
*   `stage` - the format of the exported data: `csv`, `tab`, `pipe` (default)
*   `header` - include a header (`YES`) or not (`NO`). Optional, default is `NO`.
*   `duration` - `Y` or `N` (default `N`). Includes the duration of the recording in the output (in seconds), before the location.

#### Example URL
```http
http://server/vicidial/non_agent_api.php?source=test&function=recording_lookup&stage=pipe&user=6666&pass=1234&agent_user=1000&date=2010-12-03
```

#### Example Responses
```text
ERROR: recording_lookup USER DOES NOT HAVE PERMISSION TO GET RECORDING INFO - 6666|0
ERROR: recording_lookup INVALID SEARCH PARAMETERS - 6666|1000||2010-12-03|
ERROR: recording_lookup NO RECORDINGS FOUND - 1255|6666||2010-12-03|
```

##### Successful Output Format (Pipe Delimited)
| start_time | user | recording_id | lead_id | location |
| :--- | :--- | :--- | :--- | :--- |
| 2010-12-03 12:00:01 | 1000 | 534820 | 876409 | `http://server/path/to/recording/20101203_120000_1234567890_1000-all.wav` |

##### Successful Output Format with Duration Enabled
| start_time | user | recording_id | lead_id | duration | location |
| :--- | :--- | :--- | :--- | :--- | :--- |
| 2010-12-03 12:00:01 | 1000 | 534820 | 876409 | 63 | `http://server/path/to/recording/20101203_120000_1234567890_1000-all.wav` |

---

### did_log_export
Exports all calls inbound to a DID for one day.

> [!NOTE]
> API user for this function must have `user_level` set to 7 or higher and "view reports" enabled.
> There is a hard limit of 100,000 results.

#### Required Fields
*   `phone_number` - 2-20 characters, the DID to pull logs for
*   `date` - date of the calls to pull (must be in YYYY-MM-DD format)

#### Settings Fields
*   `stage` - the format of the exported data: `csv`, `tab`, `pipe` (default)
*   `header` - include a header (`YES`) or not (`NO`). Optional, default is `NO`.

#### Example URL
```http
http://server/vicidial/non_agent_api.php?source=test&function=did_log_export&stage=pipe&user=6666&pass=1234&phone_number=3125551212&date=2010-12-03
```

#### Example Responses
```text
ERROR: did_log_export USER DOES NOT HAVE PERMISSION TO GET DID INFO - 6666|0
ERROR: did_log_export INVALID SEARCH PARAMETERS - 6666|3125551212|2010-12-03
ERROR: did_log_export NO RECORDS FOUND - 1255|3125551212|2010-12-03
```

##### Successful Output Format (Pipe Delimited)
| did_number | call_date | caller_id_number | length_in_sec |
| :--- | :--- | :--- | :--- |
| 3125551212 | 2010-12-03 12:00:01 | 7275551212 | 123 |

---

### phone_number_log
Exports list of calls placed to one or more phone numbers.

> [!NOTE]
> API user for this function must have `user_level` set to 7 or higher and "view reports" enabled.
> There is a hard limit of 100,000 results.
> *The `phone_number` field is not indexed by default. If you plan to use this function, your system administrator may need to add indexes on that field in the `vicidial_log` and `vicidial_closer_log` tables.*

#### Required Fields
*   `phone_number` - the phone number(s) to pull logs for. Allows multiple separated by commas.

#### Settings Fields
*   `stage` - the format of the exported data: `csv`, `tab`, `pipe` (default)
*   `header` - include a header (`YES`) or not (`NO`). Optional, default is `NO`.
*   `detail` - `ALL` calls or only `LAST` call. Default is `ALL`.
*   `type` - `IN` (inbound), `OUT` (outbound), or `ALL` calls. Default is `OUT`.
*   `archived_lead` - `Y` or `N` (default `N`). Query only the `vicidial_list_archive` table for lead details.

#### Example URL
```http
http://server/vicidial/non_agent_api.php?source=test&function=phone_number_log&stage=pipe&user=6666&pass=1234&phone_number=3125551212,9998887112
```

#### Example Responses
```text
ERROR: phone_number_log USER DOES NOT HAVE PERMISSION TO GET CALL LOG INFO - 6666|0
ERROR: phone_number_log NO VALID PHONE NUMBERS DEFINED - 6666|312|2010-12-03
ERROR: phone_number_log NO RECORDS FOUND - 1255|3125551212|2010-12-03
```

##### Successful Output Format (Pipe Delimited)
| phone_number | call_date | list_id | length_in_sec | lead_status | hangup_reason | call_status | source_id | user |
| :--- | :--- | :--- | :--- | :--- | :--- | :--- | :--- | :--- |
| 9998887112 | 2017-01-06 08:11:01 | 82106 | 25 | TD | AGENT | SALE | mail098 | 6666 |

---

### agent_stats_export
Exports agent activity statistics.

> [!NOTE]
> API user for this function must have `user_level` set to 7 or higher and "view reports" enabled.
> There is a hard limit of 10,000,000 records analyzed.

#### Required Fields
*   `source` - description of what originated the API call (maximum 20 characters)
*   `datetime_start` - start date/time of the agent activity to pull ("YYYY-MM-DD+HH:MM:SS" format)
*   `datetime_end` - end date/time of the agent activity to pull ("YYYY-MM-DD+HH:MM:SS" format)

#### Settings Fields
*   `agent_user` - 2-20 characters, limit to one agent <optional>
*   `campaign_id` - 2-8 characters, limit to one campaign <optional>
*   `stage` - the format of the exported data: `csv`, `tab`, `pipe` (default)
*   `header` - include a header (`YES`) or not (`NO`). Optional, default is `NO`.
*   `time_format` - time format (`H`, `HF`, `M`, `S`) in hours, minutes, or seconds. `H` = 1:23:45, `M` = 83:45, `S` = 5023 (default `HF`).
    *   *HF will force hour format even for zero seconds time "0:00:00"*
*   `group_by_campaign` - divide and output the agent results grouped by campaign (`YES`) or not (`NO`). Optional, default is `NO`.

#### Example URL
```http
http://server/vicidial/non_agent_api.php?source=test&function=agent_stats_export&time_format=M&stage=pipe&user=6666&pass=1234&agent_user=6666&datetime_start=2012-08-09+00:00:00&datetime_end=2012-08-10+23:59:59
http://server/vicidial/non_agent_api.php?DB=0&stage=pipe&user=6666&pass=1234&source=test&function=agent_stats_export&datetime_start=2012-03-26+00:00:00&datetime_end=2012-03-26+23:59:59&header=YES
```

#### Example Responses
```text
ERROR: agent_stats_export USER DOES NOT HAVE PERMISSION TO GET AGENT INFO - 6666|0
ERROR: agent_stats_export INVALID SEARCH PARAMETERS - 6666||2010-12-03 00:00:00|2010-12-03 00:00:01|
ERROR: agent_stats_export NO RECORDS FOUND - 1255||2010-12-03 00:00:00|2010-12-03 00:00:01|
```

##### Successful Output Format (Pipe Delimited)
| user | full_name | user_group | calls | login_time | total_talk_time | avg_talk_time | avg_wait_time | pct_of_queue | pause_time | sessions | avg_session | pauses | avg_pause_time | pause_pct | pauses_per_session | wait_time | talk_time | dispo_time | dead_time |
| :--- | :--- | :--- | :--- | :--- | :--- | :--- | :--- | :--- | :--- | :--- | :--- | :--- | :--- | :--- | :--- | :--- | :--- | :--- | :--- |
| 1029 | test agent | ADMIN | 1 | 0:01:52 | 0:00:16 | 0:00:16 | 0:00:18 | 16.7% | 54 | 3 | 0:00:37 | 5 | 0:00:11 | 48.2% | 2 | 0:00:32 | 0:07:54 | 0:01:22 | 0:00:00 |
| 6666 | Admin user | ADMIN | 16 | 0:34:31 | 0:07:54 | 0:00:00 | 0:00:07 | 100% | 1411 | 13 | 0:02:39 | 42 | 0:00:34 | 68.1% | 3 | 0:01:44 | 0:07:54 | 0:01:22 | 0:00:00 |

> [!NOTE]
> If the `group_by_campaign` option is enabled, the format above will have `campaign_id` as the first column, and each agent's activity will be divided on separate lines by the campaigns they were logged into during the requested time period.

---

### user_group_status
Real-time status of one or more user groups.

> [!NOTE]
> API user for this function must have `user_level` set to 7 or higher and "view reports" enabled.
> There is a hard limit of 10,000,000 records analyzed.

#### Required Fields
*   `source` - description of what originated the API call (maximum 20 characters)
*   `user_groups` - pipe-delimited list of user groups to get status information for (e.g., `ADMIN\|AGENTS`)

#### Settings Fields
*   `stage` - the format of the exported data: `csv`, `tab`, `pipe` (default)
*   `header` - include a header (`YES`) or not (`NO`). Optional, default is `NO`.

#### Example URL
```http
http://server/vicidial/non_agent_api.php?source=test&user=6666&pass=1234&function=user_group_status&user_groups=ADMIN|AGENT&header=YES
```

#### Example Responses
```text
ERROR: user_group_status USER DOES NOT HAVE PERMISSION TO GET USER GROUP INFO - 6666|0
ERROR: user_group_status INVALID SEARCH PARAMETERS - 6666||
```

##### Successful Output Format (Header Enabled)
| usergroups | calls_waiting | agents_logged_in | agents_in_calls | agents_waiting | agents_paused | agents_in_dead_calls | agents_in_dispo | agents_in_dial |
| :--- | :--- | :--- | :--- | :--- | :--- | :--- | :--- | :--- |
| ADMIN AGENT | 0 | 1 | 0 | 1 | 0 | 0 | 0 | 0 |

---

### in_group_status
Real-time status of one or more in groups.

> [!NOTE]
> API user for this function must have `user_level` set to 7 or higher and "view reports" enabled.
> There is a hard limit of 10,000,000 records analyzed.

#### Required Fields
*   `source` - description of what originated the API call (maximum 20 characters)
*   `in_groups` - pipe-delimited list of inbound groups to get status information for (e.g., `SALESLINE\|SUPPORT`)

#### Settings Fields
*   `stage` - the format of the exported data: `csv`, `tab`, `pipe` (default)
*   `header` - include a header (`YES`) or not (`NO`). Optional, default is `NO`.

#### Example URL
```http
http://server/vicidial/non_agent_api.php?source=test&user=6666&pass=1234&stage=csv&function=in_group_status&in_groups=SALESLINE|SUPPORT&header=YES
```

#### Example Responses
```text
ERROR: in_group_status USER DOES NOT HAVE PERMISSION TO GET IN-GROUP INFO - 6666|0
ERROR: in_group_status INVALID SEARCH PARAMETERS - 6666||
```

##### Successful Output Format (CSV)
| ingroups | total_calls | calls_waiting | agents_logged_in | agents_in_calls | agents_waiting | agents_paused | agents_in_dispo | agents_in_dial |
| :--- | :--- | :--- | :--- | :--- | :--- | :--- | :--- | :--- |
| SALESLINE SUPPORT | 31 | 0 | 75 | 54 | 0 | 20 | 1 | 0 |

---

### agent_status
Real-time status of one agent user.

> [!NOTE]
> API user for this function must have `user_level` set to 7 or higher and "view reports" enabled.

#### Required Fields
*   `source` - description of what originated the API call (maximum 20 characters)
*   `agent_user` - 2-20 characters, used to specify a single agent's status

#### Settings Fields
*   `stage` - the format of the exported data: `csv`, `tab`, `pipe` (default)
*   `header` - include a header (`YES`) or not (`NO`). Optional, default is `NO`.
*   `include_ip` - include computer IP field (`YES`) or not (`NO`). This is the last IP address used by the agent to log into the agent screen.

#### Example URL
```http
http://server/vicidial/non_agent_api.php?source=test&user=6666&pass=1234&function=agent_status&agent_user=1234&stage=csv&header=YES&include_ip=YES
```

#### Example Responses
```text
ERROR: agent_status USER DOES NOT HAVE PERMISSION TO GET AGENT INFO - 6666|0
ERROR: agent_status INVALID SEARCH PARAMETERS - 6666||
ERROR: agent_status AGENT NOT FOUND - 6666||
ERROR: agent_status AGENT NOT LOGGED IN - 6666||
```

##### Successful Output Format (CSV)
| status | call_id | lead_id | campaign_id | calls_today | full_name | user_group | user_level | pause_code | real_time_sub_status | phone_number | vendor_lead_code | session_id | computer_ip |
| :--- | :--- | :--- | :--- | :--- | :--- | :--- | :--- | :--- | :--- | :--- | :--- | :--- | :--- |
| INCALL | M4050908070000012345 | 12345 | TESTCAMP | 1 | Test Agent | AGENTS | 3 | LOGIN | | 7275551212 | 123456 | 8600051 | 192.168.1.205 |
| INCALL | M4181606420000000104 | 104 | TESTBLND | 1 | Admin | ADMIN | 9 | BRK2 | DEAD | 3125551212 | 123457 | 8600051 | 192.168.1.206 |
| PAUSED | | 105 | TESTBLND | 1 | Admin | ADMIN | 9 | | PREVIEW | 9545551212 | 123458 | 8600052 | 192.168.1.202 |

> [!NOTE]
> The `real_time_sub_status` field can consist of: `DEAD`, `DISPO`, `3-WAY`, `PARK`, `RING`, `PREVIEW`, `DIAL` or it can be empty.

---

### user_details
Display information about one user.

> [!NOTE]
> API user for this function must have `user_level` set to 7 or higher and "view reports" enabled.

#### Required Fields
*   `source` - description of what originated the API call (maximum 20 characters)
*   `agent_user` - 2-20 characters, specifies the user ID

#### Settings Fields
*   `stage` - the format of the exported data: `csv`, `tab`, `pipe` (default)
*   `header` - include a header (`YES`) or not (`NO`). Optional, default is `NO`.

#### Example URL
```http
http://server/vicidial/non_agent_api.php?source=test&user=6666&pass=1234&function=user_details&agent_user=1234&stage=csv&header=YES
```

#### Example Responses
```text
ERROR: user_details USER DOES NOT HAVE PERMISSION TO GET USER DETAILS - 6666|0
ERROR: user_details INVALID SEARCH PARAMETERS - 6666||
ERROR: user_details USER NOT FOUND - 6666||
```

##### Successful Output Format (CSV)
| user | full_name | user_group | user_level | active |
| :--- | :--- | :--- | :--- | :--- |
| 11111 | test user | ADMIN | 6 | Y |

---

### callid_info
Information about a call based upon the caller_code or call ID.

> [!NOTE]
> API user for this function must have `user_level` set to 7 or higher and "view reports" enabled.

#### Required Fields
*   `source` - description of what originated the API call (maximum 20 characters)
*   `call_id` - 16-40 characters

#### Settings Fields
*   `stage` - the format of the exported data: `csv`, `tab`, `pipe` (default)
*   `header` - include a header (`YES`) or not (`NO`). Optional, default is `NO`.
*   `detail` - if set to `YES`, more call info will be output. Default is `NO` (only callid and customer talk time will be output).

#### Example URL
```http
http://server/vicidial/non_agent_api.php?source=test&user=6666&pass=1234&function=callid_info&call_id=M4050908070000012345&stage=csv&header=YES
```

#### Example Responses
```text
ERROR: callid_info USER DOES NOT HAVE PERMISSION TO GET CALL INFO - 6666|0
ERROR: callid_info INVALID SEARCH PARAMETERS - 6666||
ERROR: callid_info CALL NOT FOUND - 6666||
ERROR: callid_info CALL LOG NOT FOUND - 6666||
```

##### Successful Output Format (Standard)
| call_id | custtime |
| :--- | :--- |
| M4050908070000012345 | 725 |

##### Successful Output Format (with `detail=YES`)
| call_id | custtime | call_date | campaign_id | list_id | status | user | phone |
| :--- | :--- | :--- | :--- | :--- | :--- | :--- | :--- |
| M4050908070000012345 | 725 | 2014-01-24 09:26:09 | TESTCAMP | 999 | A | 6666 | 3125551212 |

---

### lead_field_info
Pulls the value of one field of a lead.

> [!NOTE]
> API user for this function must have `user_level` set to 7 or higher and "modify leads" enabled.

#### Required Fields
*   `source` - description of what originated the API call (maximum 20 characters)
*   `lead_id` - 1-10 digits
*   `field_name` - name of lead field to pull

#### Optional Fields
*   `custom_fields` - `Y` or `N` (default `N`). Specifies if the field to pull is a custom field.
*   `list_id` - override for the `entry_list_id` of a custom field
*   `archived_lead` - `Y` or `N` (default `N`). Setting this to `Y` will query only the `vicidial_list_archive` table.

#### Example URL
```http
http://server/vicidial/non_agent_api.php?source=test&user=6666&pass=1234&function=lead_field_info&field_name=cc_value&lead_id=1139739&custom_fields=Y&list_id=999888999777
```

#### Example Responses
```text
ERROR: lead_field_info USER DOES NOT HAVE PERMISSION TO GET LEAD INFO - 6666|0
ERROR: lead_field_info INVALID SEARCH PARAMETERS - 6666||
ERROR: lead_field_info LIST NOT FOUND - 6666||
ERROR: lead_field_info LEAD NOT FOUND - 6666||
ERROR: lead_field_info LEAD FIELD NOT FOUND - 6666||
```

##### Successful Output
```text
1234567890
```

---

### lead_all_info
Outputs all lead data for a lead (optionally including custom fields).

> [!NOTE]
> API user for this function must have `user_level` set to 7 or higher, "view reports" enabled, and `modify_leads` set to 1-5.

#### Required Fields
*   `source` - description of what originated the API call (maximum 20 characters)
*   `lead_id` - 1-10 digits

#### Optional Fields
*   `phone_number` - 6-19 digits, can be used in place of a `lead_id` (may return multiple records up to 1000, ordered with the most recently added first)
*   `custom_fields` - `Y` or `N` (default `N`). Defines whether custom fields are pulled.
*   `force_entry_list_id` - 3-12 digits, overrides the lead's assigned `entry_list_id` when checking custom fields data from a different list
*   `archived_lead` - `Y` or `N` (default `N`). Query only the `vicidial_list_archive` table for details.

#### Settings Fields
*   `stage` - the format of the exported data: `csv`, `tab`, `newline`, `pipe` (default)
*   `header` - include a header (`YES`) or not (`NO`). Optional, default is `NO`.

#### Example URL
```http
http://server/vicidial/non_agent_api.php?source=test&user=6666&pass=1234&function=lead_all_info&lead_id=1079425
http://server/vicidial/non_agent_api.php?source=test&user=6666&pass=1234&function=lead_all_info&phone_number=3125551212
http://server/vicidial/non_agent_api.php?source=test&user=6666&pass=1234&function=lead_all_info&lead_id=1079425&header=YES&custom_fields=Y
http://server/vicidial/non_agent_api.php?source=test&user=6666&pass=1234&function=lead_all_info&lead_id=1079425&header=YES&custom_fields=Y&force_entry_list_id=999880
```

#### Example Responses
```text
ERROR: lead_all_info USER DOES NOT HAVE PERMISSION TO GET LEAD INFO - 6666|0
ERROR: lead_all_info LIST NOT FOUND - 6666|||
ERROR: lead_all_info LEAD NOT FOUND - 6666|||
```

##### Successful Output Format (Pipe Delimited)
```text
status|user|vendor_lead_code|source_id|list_id|gmt_offset_now|phone_code|phone_number|title|first_name|middle_initial|last_name|address1|address2|address3|city|state|province|postal_code|country_code|gender|date_of_birth|alt_phone|email|security_phrase|comments|called_count|last_local_call_time|rank|owner|entry_list_id|lead_id
TD|6666|554443||501|-4.00|1|9998887112|Mr|Billy'Oj|U|O'Reillyo|ad1|ad2|ad3|city next|FL|U|33777||U|0000-00-00|bobo'brien|test@test.com||2017-06-01 18:38:41|79|2017-06-09 14:53:27|0|4646|0|98234789
```

---

### lead_status_search
Displays all field values of all leads that match the status and call date in request.

> [!NOTE]
> API user for this function must have `user_level` set to 7 or higher and "modify leads" enabled.
> Both inbound and outbound call logs are searched; no more than 2,000 responses are displayed.

#### Required Fields
*   `source` - description of what originated the API call (maximum 20 characters)
*   `status` - status of leads to be pulled
*   `date` - called date of the leads to pull (YYYY-MM-DD format)

#### Optional Fields
*   `lead_id` - 1-10 digits, only used if `status` and `date` are empty
*   `custom_fields` - `Y` or `N` (default `N`). Display custom fields values or not.
*   `list_id` - override for the `entry_list_id` of a custom field
*   `archived_lead` - `Y` or `N` (default `N`). Query only the `vicidial_list_archive` table for details.

#### Example URL
```http
http://server/vicidial/non_agent_api.php?source=test&user=6666&pass=1234&function=lead_status_search&status=NP&date=2017-12-29&custom_fields=Y
```

#### Example Responses
```text
ERROR: lead_status_search USER DOES NOT HAVE PERMISSION TO GET LEAD INFO - 6666|0
ERROR: lead_status_search INVALID SEARCH PARAMETERS - 6666|||
ERROR: lead_status_search NO RESULTS FOUND - 6666|||
```

##### Successful Output Format
```text
;---------- START OF RECORD 1 ----------
lead_id => 823602
status => TD
user => 6666
vendor_lead_code => 1028405
source_id => I
list_id => 107
gmt_offset_now => -5.00
phone_code => 1
phone_number => 9999028405
...
test_custom_field => new value
```

---

### lead_search
Searches for leads in the `vicidial_list` table by phone numbers.

> [!NOTE]
> API user for this function must have `user_level` set to 7 or higher, "view reports" enabled, and `modify_leads` set to 1-5.
> * If no results are found for any of the searched phone numbers, an error will be displayed.
> * If matches are found, every searched phone number will have one line in the output.
> * If more than 1 matching lead is found for a phone number, the matching `lead_ids` will be separated by a dash (`-`).

#### Required Fields
*   `phone_number` - 6-19 digits (multiple separated by a comma, e.g.: `3125551212,9055551212,4075551212`)
*   `source` - description of what originated the API call (maximum 20 characters)

#### Settings Fields
*   `records` - number of records to display if multiple matches found (defaults to `1000` [most recently inserted leads shown first])
*   `archived_lead` - `Y` or `N` (default `N`). Query only the `vicidial_list_archive` table.
*   `header` - include a header (`YES`) or not (`NO`). Optional, default is `NO`.

#### Example URL
```http
http://server/vicidial/non_agent_api.php?source=test&user=6666&pass=1234&function=lead_search&phone_number=9998885112,9998887112,9998886112,4075551212
http://server/vicidial/non_agent_api.php?source=test&user=6666&pass=1234&function=lead_search&phone_number=9998885112,9055551212,9998886112,4075551212&header=YES
```

#### Example Responses
```text
ERROR: lead_search USER DOES NOT HAVE PERMISSION TO SEARCH FOR LEADS - 6666|0
ERROR: lead_search NO VALID SEARCH METHOD - ||||
ERROR: lead_search NO MATCHES FOUND IN THE SYSTEM - ||||
```

##### Successful Output Format (Pipe Delimited)
| search_key | records_found | lead_ids |
| :--- | :--- | :--- |
| 9998885112 | 4 | 1949076-1946517-1943958-1941399 |
| 9998887112 | 8 | 1946941-1946940-1946939-1946875-1946805-1946656-1946655-1944382 |
| 9998886112 | 1 | 1946891 |
| 4075551212 | 0 | |

---

### ccc_lead_info
Outputs lead data for cross-cluster-communication calls.

> [!NOTE]
> API user for this function must have `user_level` set to 7 or higher and "view reports" enabled.

#### Required Fields
*   `source` - description of what originated the API call (maximum 20 characters)
*   `call_id` - 16-40 characters

#### Settings Fields
*   `stage` - the format of the exported data: `csv`, `tab`, `newline`, `pipe` (default)
*   `header` - include a header (`YES`) or not (`NO`). Optional, default is `NO`.

#### Example URL
```http
http://server/vicidial/non_agent_api.php?source=test&user=6666&pass=1234&function=ccc_lead_info&call_id=M4050908070000012345&stage=csv&header=YES
```

#### Example Responses
```text
ERROR: ccc_lead_info USER DOES NOT HAVE PERMISSION TO GET LEAD INFO - 6666|0
ERROR: ccc_lead_info INVALID SEARCH PARAMETERS - 6666||
ERROR: ccc_lead_info CALL NOT FOUND - 6666||
ERROR: ccc_lead_info LEAD NOT FOUND - 6666||
```

##### Successful Output Format (Pipe Delimited)
```text
status|user|vendor_lead_code|source_id|list_id|gmt_offset_now|phone_code|phone_number|title|first_name|middle_initial|last_name|address1|address2|address3|city|state|province|postal_code|country_code|gender|date_of_birth|alt_phone|email|security_phrase|comments|called_count|last_local_call_time|rank|owner
TD|6666|554443||501|-4.00|1|9998887112|Mr|Billy'Oj|U|O'Reillyo|ad1|ad2|ad3|city next|FL|U|33777||U|0000-00-00|bobo'brien|test@test.com||2017-06-01 18:38:41|79|2017-06-09 14:53:27|0|4646
```

---

### lead_callback_info
Outputs scheduled callback data for a specific lead.

> [!NOTE]
> API user for this function must have `user_level` set to 7 or higher and "view reports" enabled.

#### Required Fields
*   `source` - description of what originated the API call (maximum 20 characters)
*   `lead_id` - 16-40 characters

#### Settings Fields
*   `stage` - the format of the exported data: `csv`, `tab`, `newline`, `pipe` (default)
*   `header` - include a header (`YES`) or not (`NO`). Optional, default is `NO`.
*   `search_location` - Where to check for scheduled callbacks (default is `CURRENT`):
    *   `CURRENT` - check for active scheduled callbacks
    *   `ARCHIVE` - check for archived scheduled callbacks
    *   `ALL` - check for both active and archived scheduled callbacks

#### Example URL
```http
http://server/vicidial/non_agent_api.php?source=test&user=6666&pass=1234&function=lead_callback_info&lead_id=12345&stage=pipe&header=YES&search_location=ALL
http://server/vicidial/non_agent_api.php?source=test&user=6666&pass=1234&function=lead_callback_info&lead_id=846446&search_location=ALL&header=YES
```

#### Example Responses
```text
ERROR: lead_callback_info USER DOES NOT HAVE PERMISSION TO GET LEAD INFO - 6666|0
ERROR: lead_callback_info INVALID SEARCH PARAMETERS - 6666||
ERROR: lead_callback_info LEAD NOT FOUND - 6666||
ERROR: lead_callback_info CALLBACK NOT FOUND - 6666||
```

##### Successful Output Format (Header Enabled)
*(The most recent callback entry will be at the bottom of the results)*

| lead_id | callback_type | recipient | callback_status | lead_status | callback_date | entry_date | user | list_id | campaign_id | user_group | customer_timezone | customer_time | comments |
| :--- | :--- | :--- | :--- | :--- | :--- | :--- | :--- | :--- | :--- | :--- | :--- | :--- | :--- |
| 846446 | ARCHIVE | USERONLY | INACTIVE | CBTEST | 2018-09-01 15:00:00 | 2018-08-27 01:06:45 | 6666 | 107 | TESTCAMP | ADMIN | AUS-ACST-N-Northern Territory Time Zone | 2018-08-27 00:00:00 | |
| 846446 | ARCHIVE | USERONLY | INACTIVE | CBTEST | 2018-12-14 14:00:00 | 2018-12-14 16:42:32 | 6666 | 107 | TESTCAMP | ADMIN | | 2018-12-14 14:00:00 | Call back after lunch |
| 846446 | CURRENT | USERONLY | INACTIVE | CBTEST | 2018-12-14 19:00:00 | 2018-12-14 16:44:08 | 6666 | 107 | TESTCAMP | ADMIN | | 2018-12-14 19:00:00 | Call back after dinner |
| 846446 | CURRENT | USERONLY | ACTIVE | CBLEAD | 2018-12-17 10:00:00 | 2018-12-14 16:50:56 | 6666 | 107 | TESTCAMP | ADMIN | USA-EST-Y-Eastern US Time Zone | 2018-12-17 10:00:00 | try back Monday morning |

---

### lead_dearchive
Moves a lead from the archive to active leads table.

> [!NOTE]
> API user for this function must have `user_level` set to 8 or higher and "modify leads" enabled.

#### Required Fields
*   `source` - description of what originated the API call (maximum 20 characters)
*   `lead_id` - 16-40 characters

#### Example URL
```http
http://server/vicidial/non_agent_api.php?source=test&user=6666&pass=1234&function=lead_dearchive&lead_id=274071
```

#### Example Responses
```text
ERROR: lead_dearchive USER DOES NOT HAVE PERMISSION TO MODIFY LEAD INFO - 6666|0
ERROR: lead_dearchive INVALID SEARCH PARAMETERS - 6666||
ERROR: lead_dearchive LEAD INSERT FAILED - 6666||
ERROR: lead_dearchive NON-ARCHIVED LEAD FOUND - 6666||
ERROR: lead_dearchive ARCHIVED LEAD NOT FOUND - 6666||
SUCCESS: lead_dearchive LEAD DE-ARCHIVED - 6666|274071|1|1
```

---

### update_log_entry
Updates the status of a `vicidial_log` or `vicidial_closer_log` entry.

> [!NOTE]
> API user for this function must have `user_level` set to 8 or higher and "modify leads" enabled.
> There is a hard limit of 10,000,000 records analyzed.

#### Required Fields
*   `source` - description of what originated the API call (maximum 20 characters)
*   `call_id` - either the uniqueid or the 20-character ID of the call
*   `group` - either a `campaign_id` or an in-group `group_id`
*   `status` - the new status for the log record to be set to

#### Example URL
```http
http://server/vicidial/non_agent_api.php?source=test&user=6666&pass=1234&function=update_log_entry&group=SALESLINE&call_id=Y2010309520000000234&status=SALEX
```

#### Example Responses
```text
ERROR: update_log_entry USER DOES NOT HAVE PERMISSION TO UPDATE LOGS - 6666|0
ERROR: update_log_entry INVALID SEARCH PARAMETERS - 6666||
ERROR: update_log_entry GROUP NOT FOUND - 6666||
ERROR: update_log_entry NO RECORDS FOUND - 6666||||
ERROR: update_log_entry NO RECORDS UPDATED - 6666|TESTCAMP|M9131554060000426264|SALEX|SALEX|1347566056.399
SUCCESS: update_log_entry RECORD HAS BEEN UPDATED - 6666|6666|TESTCAMP|M9131554060000426264|SALEX|WN|1347566056.399|
```

---

### add_lead
Adds a new lead to the `vicidial_list` table with several fields and options.

> [!NOTE]
> API user for this function must have `modify_leads` set to 1 and `user_level` set to 8 or higher.
> For fields with spaces in values, replace the space with a plus (`+`) sign.

#### Required Fields
*   `phone_number` - must be all numbers, 6-16 digits
*   `phone_code` - must be all numbers, 1-4 digits (defaults to `1` if not set)
*   `list_id` - must be all numbers, 3-12 digits (defaults to `999` if not set)
*   `source` - description of what originated the API call (maximum 20 characters)

#### Settings Fields
*   `dnc_check` - `Y`, `N` or `AREACODE`, default is `N`
*   `campaign_dnc_check` - `Y`, `N` or `AREACODE`, default is `N`
*   `campaign_id` - 2-8 Character campaign ID, required if using campaign DNC check or callbacks
*   `add_to_hopper` - `Y` or `N` (default `N`)
*   `hopper_priority` - `99` to `-99` (higher is prioritized, default is `0`)
*   `hopper_local_call_time_check` - `Y` or `N` (default `N`). Validate local/state call times before insertion.
*   `duplicate_check` - Check for duplicates. Can combine multiple options (e.g. `duplicate_check=DUPLIST-DUPTITLEALTPHONELIST`). Options:
    *   `DUPLIST` - duplicate phone number in same list
    *   `DUPCAMP` - duplicate phone number in all lists for this campaign
    *   `DUPSYS` - duplicate phone number in entire system
    *   `DUPPHONEALTLIST` - duplicate phone against phone and alt_phone in same list
    *   `DUPPHONEALTCAMP` - duplicate phone against phone and alt_phone in campaign lists
    *   `DUPPHONEALTSYS` - duplicate phone against phone and alt_phone in entire system
    *   `DUPTITLEALTPHONELIST` - duplicate title and alt_phone in same list
    *   `DUPTITLEALTPHONECAMP` - duplicate title and alt_phone in campaign lists
    *   `DUPTITLEALTPHONESYS` - duplicate title and alt_phone in entire system
    *   `DUPNAMEPHONELIST` - duplicate name and phone number in same list
    *   `DUPNAMEPHONECAMP` - duplicate name and phone number in campaign lists
    *   `DUPNAMEPHONESYS` - duplicate name and phone number in entire system
    *   Append `1/2/3/7/14/15/21/28/30/60/90/180/360DAY` to limit search range (e.g. `DUPSYS90DAY`).
*   `usacan_prefix_check` - `Y` or `N` (default `N`). Check for valid 4th digit for US/Canada (cannot be 0 or 1).
*   `usacan_areacode_check` - `Y` or `N` (default `N`). Check for valid US/Canada areacode (also checks 10-digit length).
*   `nanpa_ac_prefix_check` - `Y` or `N` (default `N`). Check valid NANPA areacode and prefix if loaded.
*   `custom_fields` - `Y` or `N` (default `N`). Accepts custom field data. To populate, add the field label as a variable to the URL (e.g. `&favorite_color=blue`).
*   `tz_method` - `<empty>`, `POSTAL`, `TZCODE` or `NANPA` (default `<empty>`, uses country code/areacode)
*   `callback` - `Y` or `N` (default `N`). Sets lead as a scheduled callback. `campaign_id` is required.
*   `callback_status` - 1-6 Characters, default is `CALLBK`
*   `callback_datetime` - YYYY-MM-DD+HH:MM:SS. 'NOW' or 'xDAYS' (where 'x' is days in future) can be used.
*   `callback_type` - `USERONLY` or `ANYONE` (default `ANYONE`)
*   `callback_user` - User ID for USERONLY callback
*   `callback_comments` - Optional comments
*   `lookup_state` - `Y` or `N` (default `N`). Looks up state from areacode if state is empty.
*   `list_exists_check` - `Y` or `N` (default `N`). Returns ERROR if the `list_id` is not defined in the system.

#### Optional Fields
*   `vendor_lead_code` - 1-20 characters
*   `source_id` - 1-50 characters
*   `gmt_offset_now` - overridden by auto-lookup of phone/areacode
*   `title` - 1-4 characters
*   `first_name` - 1-30 characters
*   `middle_initial` - 1 character
*   `last_name` - 1-30 characters
*   `address1` - 1-100 characters
*   `address2` - 1-100 characters
*   `address3` - 1-100 characters
*   `city` - 1-50 characters
*   `state` - 2 characters
*   `province` - 1-50 characters
*   `postal_code` - 1-10 characters
*   `country_code` - 3 characters
*   `gender` - `U`, `M`, `F` (default `U`)
*   `date_of_birth` - YYYY-MM-DD
*   `alt_phone` - 1-12 characters
*   `email` - 1-70 characters
*   `security_phrase` - 1-100 characters
*   `comments` - 1-255 characters
*   `multi_alt_phones` - 5-1024 characters
*   `rank` - 1-5 digits
*   `owner` - 1-20 characters
*   `entry_list_id` - *Warning: Can break custom fields. Do not use if `custom_fields` is set to `Y`.*

##### Multi-ALT-Phones Format
```text
7275551212_1_work!7275551213_1_sister+house!1234567890_1_neighbor
```
Format: `phone-number`\_`phone-code`\_`phone-note`, separated by exclamation marks `!`.

#### Example URLs
```http
http://server/vicidial/non_agent_api.php?source=test&user=6666&pass=1234&function=add_lead&phone_number=7275551111

http://server/vicidial/non_agent_api.php?source=test&user=6666&pass=1234&function=add_lead&phone_number=7275551212&phone_code=1&list_id=999&dnc_check=N&first_name=Bob&last_name=Wilson

http://server/vicidial/non_agent_api.php?source=test&user=6666&pass=1234&function=add_lead&phone_number=7275551111&phone_code=1&list_id=999&dnc_check=N&first_name=Bob&last_name=Wilson&add_to_hopper=Y&hopper_local_call_time_check=Y

http://server/vicidial/non_agent_api.php?source=test&user=6666&pass=1234&function=add_lead&phone_number=7275551111&phone_code=1&list_id=999&dnc_check=N&campaign_dnc_check=Y&campaign_id=TESTCAMP&first_name=Bob&last_name=Wilson&address1=1234+Main+St.&city=Chicago+Heights&state=IL&add_to_hopper=Y&hopper_local_call_time_check=Y&multi_alt_phones=7275551212_1_work!7275551213_1_sister+house!1234567890_1_neighbor

http://server/vicidial/non_agent_api.php?source=test&user=6666&pass=1234&function=add_lead&phone_number=7275551212&phone_code=1&list_id=999&dnc_check=N&first_name=Bob&last_name=Wilson&duplicate_check=DUPPHONEALTLIST-DUPNAMEPHONELIST

http://server/vicidial/non_agent_api.php?source=test&user=6666&pass=1234&function=add_lead&phone_number=7275551212&phone_code=1&list_id=999&custom_fields=Y&favorite_color=blue

http://server/vicidial/non_agent_api.php?source=test&user=6666&pass=1234&function=add_lead&phone_number=7275551111&campaign_id=TESTCAMP&callback=Y&callback_status=CALLBK&callback_datetime=NOW&callback_type=USERONLY&callback_user=6666&callback_comments=Comments+go+here
```

#### Example Responses
```text
SUCCESS: add_lead LEAD HAS BEEN ADDED - 7275551111|6666|999|193715|-4
NOTICE: add_lead ADDED TO HOPPER - 7275551111|6666|193715|1677922

SUCCESS: add_lead LEAD HAS BEEN ADDED - 7275551111|6666|999|193716|-4
NOTICE: add_lead CUSTOM FIELDS VALUES ADDED - 7275551111|1234|101
NOTICE: add_lead CUSTOM FIELDS NOT ADDED, CUSTOM FIELDS DISABLED - 7275551111|Y|0
NOTICE: add_lead CUSTOM FIELDS NOT ADDED, NO CUSTOM FIELDS DEFINED FOR THIS LIST - 7275551111|1234|101
NOTICE: add_lead CUSTOM FIELDS NOT ADDED, NO FIELDS DEFINED - 7275551111|1234|101
NOTICE: add_lead MULTI-ALT-PHONE NUMBERS LOADED - 3|6666|193716
NOTICE: add_lead NOT ADDED TO HOPPER, OUTSIDE OF LOCAL TIME - 7275551111|6666|193716|-4|0

NOTICE: add_lead SCHEDULED CALLBACK ADDED - 1234|2011-09-29 12:00:01|TESTCAMP|6666|USERONLY|CALLBK
NOTICE: add_lead SCHEDULED CALLBACK NOT ADDED, USER NOT VALID - 1234|TESTCAMP|6|
NOTICE: add_lead SCHEDULED CALLBACK NOT ADDED, CAMPAIGN NOT VALID - 1234|XYZ

NOTICE: add_lead NANPA options disabled, NANPA prefix data not loaded - 0|6666

ERROR: add_lead INVALID PHONE NUMBER LENGTH - 72755|6666 
ERROR: add_lead INVALID PHONE NUMBER PREFIX - 72755|6666 
ERROR: add_lead INVALID PHONE NUMBER AREACODE - 72755|6666 
ERROR: add_lead INVALID PHONE NUMBER NANPA AREACODE PREFIX - 7275551212|6666

ERROR: add_lead USER DOES NOT HAVE PERMISSION TO ADD LEADS TO THE SYSTEM - 6666|0 
ERROR: add_lead NOT AN ALLOWED LIST ID - 7275551212|98762
ERROR: add_lead NOT A DEFINED LIST ID, LIST EXISTS CHECK ENABLED - 7275551212|12344

ERROR: NO FUNCTION SPECIFIED

ERROR: add_lead DUPLICATE PHONE NUMBER IN LIST - 7275551111|101|8765444
ERROR: add_lead DUPLICATE PHONE NUMBER IN CAMPAIGN LISTS - 7275551111|101|8765444|101
ERROR: add_lead DUPLICATE PHONE NUMBER IN SYSTEM - 7275551111|101|8765444|101
ERROR: add_lead DUPLICATE PHONE NUMBER IN LIST - 7275551111|101|8765444|PHONE
ERROR: add_lead DUPLICATE PHONE NUMBER IN CAMPAIGN LISTS - 7275551111|101|8765444|101|ALT
ERROR: add_lead DUPLICATE PHONE NUMBER IN SYSTEM - 7275551111|101|8765444|101|PHONE
ERROR: add_lead DUPLICATE TITLE ALT_PHONE IN LIST - 1234|7275551111|101|8765444
ERROR: add_lead DUPLICATE TITLE ALT_PHONE IN CAMPAIGN LISTS - 1234|7275551111|101|8765444|101
ERROR: add_lead DUPLICATE TITLE ALT_PHONE IN SYSTEM - 1234|7275551111|101|8765444|101
ERROR: add_lead DUPLICATE NAME PHONE IN LIST - Bob|Smith|7275551113|101|8765444|101
ERROR: add_lead DUPLICATE NAME PHONE IN CAMPAIGN LISTS - Bob|Smith|7275551113|101|8765444|101
ERROR: add_lead DUPLICATE NAME PHONE IN SYSTEM - Bob|Smith|7275551113|101|8765444|101
```

---

### update_lead
Updates lead information in the `vicidial_list` and custom fields tables.

> [!NOTE]
> API user for this function must have `modify_leads` set to 1 and `user_level` set to 8 or higher.
> * To set a field to empty, use the value `--BLANK--` (e.g. `&province=--BLANK--`).
> * Avoid special characters like apostrophes, quotes, or ampersands.

#### Required Fields
*   `lead_id` - must be all numbers, 1-9 digits (not required if using `vendor_lead_code` or `phone_number` for search)
*   `vendor_lead_code` - can be used instead of `lead_id` to match leads
*   `phone_number` - can be used instead of `lead_id` or `vendor_lead_code` to match leads
*   `source` - description of what originated the API call (maximum 20 characters)

#### Settings Fields
*   `search_method` - Combine parameters to search (`LEAD_ID`, `VENDOR_LEAD_CODE`, `PHONE_NUMBER`). E.g.: `&search_method=LEAD_ID_VENDOR_LEAD_CODE`. (Search order is fixed: Lead ID > Vendor Lead Code > Phone Number. Default is `LEAD_ID`)
*   `search_location` - Where to check for records: `LIST` (same list), `CAMPAIGN` (all lists in campaign), `SYSTEM` (entire system, default)
*   `insert_if_not_found` - `Y` or `N` (default `N`). Inserts as a new lead if no match is found (requires `phone_code`, `phone_number` and `list_id`).
*   `records` - number of records to update if multiple found (defaults to `1` [most recent lead])
*   `custom_fields` - `Y` or `N` (default `N`). To update custom fields, add the field label as a variable to the URL (e.g. `&favorite_color=blue`).
*   `no_update` - `Y` or `N` (default `N`). If set to `Y`, only checks if a matching lead exists and returns status.
*   `delete_lead` - `Y` or `N` (default `N`). Deletes the lead from the `vicidial_list` table.
*   `delete_cf_data` - `Y` or `N` (default `N`). Deletes custom fields data and resets `entry_list_id`.
*   `reset_lead` - `Y` or `N` (default `N`). Resets called-since-last-reset flag.
*   `callback` - `Y`, `N` or `REMOVE` (default `N`). Sets lead as a scheduled callback. `REMOVE` deletes callback entries.
*   `callback_status` - 1-6 Characters, default is `CALLBK`
*   `callback_datetime` - YYYY-MM-DD+HH:MM:SS. Supports 'NOW' or 'xDAYS'.
*   `callback_type` - `USERONLY` or `ANYONE` (default `ANYONE`)
*   `callback_user` - User ID for USERONLY callback
*   `callback_comments` - Optional comments
*   `update_phone_number` - `Y` or `N` (default `N`). Enable to update the phone number.
*   `add_to_hopper` - `Y` or `N` (default `N`)
*   `remove_from_hopper` - `Y` or `N` (default `N`)
*   `hopper_priority` - `99` to `-99` (default `0`)
*   `hopper_local_call_time_check` - `Y` or `N` (default `N`).
*   `list_exists_check` - `Y` or `N` (default `N`).
*   `archived_lead` - `Y` or `N` (default `N`).

#### Editable Fields
*   `user_field` - 1-20 characters, updates the 'user' field in the `vicidial_list` table
*   `list_id_field` - 3-12 digits, updates 'list_id'
*   `status` - 1-6 characters
*   `vendor_lead_code` - 1-20 characters
*   `source_id` - 1-50 characters
*   `gmt_offset_now` - overridden by auto-lookup
*   `title` - 1-4 characters
*   `first_name` - 1-30 characters
*   `middle_initial` - 1 character
*   `last_name` - 1-30 characters
*   `address1` - 1-100 characters
*   `address2` - 1-100 characters
*   `address3` - 1-100 characters
*   `city` - 1-50 characters
*   `state` - 2 characters
*   `province` - 1-50 characters
*   `postal_code` - 1-10 characters
*   `country_code` - 3 characters
*   `gender` - `U`, `M`, `F` (default `U`)
*   `date_of_birth` - YYYY-MM-DD
*   `alt_phone` - 1-12 characters
*   `email` - 1-70 characters
*   `security_phrase` - 1-100 characters
*   `comments` - 1-255 characters
*   `rank` - 1-5 digits
*   `owner` - 1-20 characters
*   `called_count` - digits only
*   `phone_code` - digits only (1-4 digits)
*   `entry_list_id` - *Warning: Can break custom fields.*
*   `force_entry_list_id` - *Warning: Forces entry_list_id override across all custom fields queries executed.*

#### Example URLs
```http
http://server/vicidial/non_agent_api.php?source=test&user=6666&pass=1234&function=update_lead&lead_id=27&last_name=SMITH

http://server/vicidial/non_agent_api.php?source=test&user=6666&pass=1234&function=update_lead&search_method=VENDOR_LEAD_CODE&vendor_lead_code=1000019&last_name=JOHNSON

http://server/vicidial/non_agent_api.php?source=test&user=6666&pass=1234&function=update_lead&search_method=PHONE_NUMBER&records=2&list_id=8107&search_location=LIST&phone_number=9999000019&last_name=WILSON

http://server/vicidial/non_agent_api.php?source=test&user=6666&pass=1234&function=update_lead&lead_id=405794&last_name=SMITH&city=Chicago&custom_fields=Y&favorite_color=blue

http://server/vicidial/non_agent_api.php?source=test&user=6666&pass=1234&function=update_lead&search_location=LIST&search_method=PHONE_NUMBER&insert_if_not_found=Y&phone_number=9999000029&phone_code=1&list_id=999&first_name=Bob&last_name=Wilson&city=Chicago&custom_fields=Y&favorite_color=red

http://server/vicidial/non_agent_api.php?source=test&user=6666&pass=1234&function=update_lead&insert_if_not_found=Y&search_method=VENDOR_LEAD_CODE_PHONE_NUMBER&vendor_lead_code=89763545&phone_number=7275551212&phone_code=1&list_id=999&first_name=Bob&last_name=Wilson&custom_fields=Y&favorite_color=blue

http://server/vicidial/non_agent_api.php?source=test&user=6666&pass=1234&function=update_lead&search_location=LIST&search_method=VENDOR_LEAD_CODE_PHONE_NUMBER&insert_if_not_found=Y&phone_number=9999000029&phone_code=1&list_id=999&first_name=Bob&last_name=Wilson&city=Chicago&custom_fields=Y&favorite_color=red&user_field=1008&list_id_field=107&status=OLD

http://server/vicidial/non_agent_api.php?source=test&user=6666&pass=1234&function=update_lead&lead_id=27&no_update=Y

http://server/vicidial/non_agent_api.php?source=test&user=6666&pass=1234&function=update_lead&lead_id=27&delete_lead=Y

http://server/vicidial/non_agent_api.php?source=test&user=6666&pass=1234&function=update_lead&lead_id=27&delete_lead=Y&custom_fields=Y

http://server/vicidial/non_agent_api.php?source=test&user=6666&pass=1234&function=update_lead&lead_id=27&delete_cf_data=Y

http://server/vicidial/non_agent_api.php?source=test&user=6666&pass=1234&function=update_lead&lead_id=406757&campaign_id=TESTCAMP&callback=Y&callback_status=CALLBK&callback_datetime=NOW&callback_type=USERONLY&callback_user=1028&callback_comments=Comments+go+here+again

http://server/vicidial/non_agent_api.php?source=test&user=6666&pass=1234&function=update_lead&search_location=SYSTEM&search_method=PHONE_NUMBER&phone_number=9998887112&no_update=Y&add_to_hopper=Y&hopper_priority=99&hopper_local_call_time_check=Y
```

#### Example Responses
```text
SUCCESS: update_lead LEAD HAS BEEN UPDATED - 6666|193716
NOTICE: update_lead CUSTOM FIELDS VALUES UPDATED - 7275551111|1234|101
NOTICE: update_lead CUSTOM FIELDS NOT UPDATED, CUSTOM FIELDS DISABLED - 7275551111|Y|0
NOTICE: update_lead CUSTOM FIELDS NOT UPDATED, NO CUSTOM FIELDS DEFINED FOR THIS LIST - 7275551111|1234|101
NOTICE: update_lead CUSTOM FIELDS NOT UPDATED, NO FIELDS DEFINED - 7275551111|1234|101

NOTICE: update_lead SCHEDULED CALLBACK UPDATED - 1234|2011-09-29 12:00:01|TESTCAMP|6666|USERONLY|CALLBK
NOTICE: update_lead SCHEDULED CALLBACK NOT UPDATED, NO FIELDS SPECIFIED - 1234|
NOTICE: update_lead SCHEDULED CALLBACK ADDED - 1234|2011-09-29 12:00:01|TESTCAMP|6666|USERONLY|CALLBK
NOTICE: update_lead SCHEDULED CALLBACK NOT ADDED, USER NOT VALID - 1234|TESTCAMP|6|
NOTICE: update_lead SCHEDULED CALLBACK NOT ADDED, CAMPAIGN NOT VALID - 1234|XYZ

NOTICE: update_lead NO MATCHES FOUND IN THE SYSTEM - 6666|4567|897654327|7275551212
SUCCESS: update_lead LEAD HAS BEEN ADDED - 7275551111|6666|999|193716|-4

SUCCESS: update_lead LEAD HAS BEEN DELETED - 7275551111|6666|999|193716|-4
NOTICE: update_lead CUSTOM FIELDS ENTRY DELETED 1 - 7275551111|6666|999|193716|-4
NOTICE: update_lead CUSTOM FIELDS DATA HAS BEEN DELETED - 6666|193716|1|1

NOTICE: update_lead LEADS FOUND IN THE SYSTEM: |6666|1010542|12345678901234|9998887112|3333444|0
                                               (user|lead_id|vendorleadcode|phone     |list_id|entry_list_id)

NOTICE: update_lead ADDED TO HOPPER - 7275551111|6666|193715|1677922
NOTICE: update_lead NOT ADDED TO HOPPER - 7275551111|6666|193715|1677922
NOTICE: update_lead NOT ADDED TO HOPPER, OUTSIDE OF LOCAL TIME - 7275551111|193715|-5|0|6666
NOTICE: update_lead NOT ADDED TO HOPPER, LEAD NOT FOUND - 7275551111|193715|6666
NOTICE: update_lead NOT ADDED TO HOPPER, LEAD IS ALREADY IN THE HOPPER - 7275551111|193715|6666
NOTICE: update_lead REMOVED FROM HOPPER - 7275551111|193715|READY|6666
NOTICE: update_lead NOT REMOVED FROM HOPPER - 7275551111|193715|DNC|6666
NOTICE: update_lead NOT REMOVED FROM HOPPER, LEAD IS NOT IN THE HOPPER - 7275551111|193715|6666

ERROR: update_lead INVALID DATA FOR LEAD INSERTION - 6666|||
ERROR: update_lead NO MATCHES FOUND IN THE SYSTEM - 6666|||
ERROR: update_lead NO VALID SEARCH METHOD - 6666|SYSTEM|||
ERROR: update_lead NOT A DEFINED LIST ID, LIST EXISTS CHECK ENABLED - 6666|139

ERROR: update_lead USER DOES NOT HAVE PERMISSION TO UPDATE LEADS IN THE SYSTEM - 6666|0 
ERROR: update_lead NOT AN ALLOWED LIST ID, SEARCH - 7275551212|98762
ERROR: update_lead NOT AN ALLOWED LIST ID, RESULTS - 7275551212|98762
ERROR: update_lead NOT AN ALLOWED LIST ID, INSERT - 7275551212|98762
```

---

### batch_update_lead
Updates multiple leads in `vicidial_list` with a single set of field values.

> [!NOTE]
> API user for this function must have `modify_leads` set to 1 and `user_level` set to 8 or higher.
> This function will first find which leads exist and then update them all at once. There will not be admin log entries associated with each individual lead modified.

#### Required Fields
*   `lead_ids` - must be all numbers and commas (1-9 digit numbers separated by a comma)
*   `source` - description of what originated the API call (maximum 20 characters)

#### Settings Fields
*   `list_exists_check` - `Y` or `N` (default `N`). Errors if the updated `list_id_field` is not defined in the system.
*   `reset_lead` - `Y` or `N` (default `N`). Resets called-since-last-reset flags.
*   `records` - maximum number of records to update (defaults to `100` [most recently created leads])
*   `archived_lead` - `Y` or `N` (default `N`).

#### Editable Fields
*   `user_field` - 1-20 characters
*   `list_id_field` - 3-12 digits
*   `status` - 1-6 characters
*   `vendor_lead_code` - 1-20 characters
*   `source_id` - 1-50 characters
*   `gmt_offset_now` - overridden by auto-lookup
*   `title` - 1-4 characters
*   `first_name` - 1-30 characters
*   `middle_initial` - 1 character
*   `last_name` - 1-30 characters
*   `address1` - 1-100 characters
*   `address2` - 1-100 characters
*   `address3` - 1-100 characters
*   `city` - 1-50 characters
*   `state` - 2 characters
*   `province` - 1-50 characters
*   `postal_code` - 1-10 characters
*   `country_code` - 3 characters
*   `gender` - `U`, `M`, `F` (default `U`)
*   `date_of_birth` - YYYY-MM-DD
*   `alt_phone` - 1-12 characters
*   `email` - 1-70 characters
*   `security_phrase` - 1-100 characters
*   `comments` - 1-255 characters
*   `rank` - 1-5 digits
*   `owner` - 1-20 characters
*   `called_count` - digits only
*   `phone_code` - digits only (1-4 digits)

#### Example URL
```http
http://server/vicidial/non_agent_api.php?source=test&user=6666&pass=1234&function=batch_update_lead&lead_ids=2345,2346&list_id_field=138
```

#### Example Responses
```text
SUCCESS: batch_update_lead LEADS HAVE BEEN UPDATED - 6666|2|2345,2346
ERROR: batch_update_lead USER DOES NOT HAVE PERMISSION TO UPDATE LEADS IN THE SYSTEM - 6666|0|
ERROR: batch_update_lead NO LEADS DEFINED - 6666|
ERROR: batch_update_lead NOT AN ALLOWED LIST ID - 6666|138
ERROR: batch_update_lead NOT A DEFINED LIST ID, LIST EXISTS CHECK ENABLED - 6666|138
ERROR: batch_update_lead NO LEADS FOUND - 6666|0|2345,2346
```

---

### add_user
Adds a user to the system.

> [!NOTE]
> API user for this function must have `user_level` set to 8 or higher and "modify users" enabled.
> This function does not work with Vtiger integration.

#### Required Fields
*   `agent_user` - 2-20 characters (for auto-generated send 'AUTOGENERATED')
*   `agent_pass` - 1-20 characters
*   `agent_user_level` - 1 through 9 (one digit only)
*   `agent_full_name` - 1-50 characters
*   `agent_user_group` - 1-20 characters, must be a valid user group

#### Optional Fields
*   `phone_login` - 1-20 characters
*   `phone_pass` - 1-20 characters
*   `hotkeys_active` - `0` or `1`
*   `voicemail_id` - 1-10 digits
*   `email` - 1-100 characters
*   `custom_one` - 1-100 characters
*   `custom_two` - 1-100 characters
*   `custom_three` - 1-100 characters
*   `custom_four` - 1-100 characters
*   `custom_five` - 1-100 characters
*   `wrapup_seconds_override` - from -1 to 9999
*   `agent_choose_ingroups` - `0` or `1` (default `1`)
*   `agent_choose_blended` - `0` or `1` (default `1`)
*   `closer_default_blended` - `0` or `1` (default `0`)
*   `in_groups` - pipe-delimited list of inbound groups (e.g. `SALESLINE\|SUPPORT`)

#### Example URL
```http
http://server/vicidial/non_agent_api.php?source=test&function=add_user&user=6666&pass=1234&agent_user=1000&agent_pass=9999&agent_user_level=1&agent_full_name=Testing+Person&agent_user_group=AGENTS
http://server/vicidial/non_agent_api.php?source=test&function=add_user&user=6666&pass=1234&agent_user=1000&agent_pass=9999&agent_user_level=1&agent_full_name=Testing+Person&agent_user_group=AGENTS&phone_login=100&phone_pass=test&hotkeys_active=1&voicemail_id=&custom_one=jjjj&custom_two=yyyy&custom_three=kkk&custom_four=wwww&custom_five=hhhh
```

#### Example Responses
```text
ERROR: add_user USER DOES NOT HAVE PERMISSION TO ADD USERS - 6666|0
ERROR: add_user YOU MUST USE ALL REQUIRED FIELDS - 6666|1000||||
ERROR: add_user USER DOES NOT HAVE PERMISSION TO ADD USERS IN THIS USER LEVEL - 6666|8
ERROR: add_user USER GROUP DOES NOT EXIST - 6666|ADFHR
ERROR: add_user USER ALREADY EXISTS - 6666|6666
SUCCESS: add_user USER HAS BEEN ADDED - 6666|1000|1234|8|Testing+Person|AGENTS
```

---

### copy_user
Copies an existing user in the system to a new user ID and name.

> [!NOTE]
> API user for this function must have `user_level` set to 8 or higher and "modify users" enabled.
> This function does not work with Vtiger integration.

#### Required Fields
*   `agent_user` - 2-20 characters (for auto-generated send 'AUTOGENERATED')
*   `agent_pass` - 1-20 characters
*   `agent_full_name` - 1-50 characters
*   `source_user` - 2-20 characters (must be an existing user ID in the system)

#### Example URL
```http
http://server/vicidial/non_agent_api.php?source=test&function=copy_user&user=6666&pass=1234&agent_user=1000&agent_pass=9999&source_user=5555&agent_full_name=Copy+Person
```

#### Example Responses
```text
ERROR: copy_user USER DOES NOT HAVE PERMISSION TO COPY USERS - 6666|0
ERROR: copy_user YOU MUST USE ALL REQUIRED FIELDS - 6666|1000|5555|||
ERROR: copy_user USER DOES NOT HAVE PERMISSION TO COPY USERS IN THIS USER LEVEL - 9|8
ERROR: copy_user SOURCE USER DOES NOT EXIST - 6666|5555
ERROR: copy_user USER GROUP DOES NOT EXIST - 6666|ADFHR
ERROR: copy_user USER ALREADY EXISTS - 6666|6666
SUCCESS: copy_user USER HAS BEEN COPIED - 6666|1000|1234|Testing+Person|5555
```

---

### update_user
Updates or deletes a user entry already in the system.

> [!NOTE]
> API user for this function must have `user_level` set to 8 or higher and "modify user" enabled.

#### Required Fields
*   `agent_user` - 2-20 characters

#### Optional Fields
*   `agent_pass` - 1-20 characters
*   `agent_user_level` - 1 through 9 (one digit only)
*   `agent_full_name` - 1-50 characters
*   `agent_user_group` - 1-20 characters, must be a valid user group
*   `phone_login` - 1-20 characters
*   `phone_pass` - 1-20 characters
*   `hotkeys_active` - `0` or `1`
*   `voicemail_id` - 1-10 digits
*   `email` - 1-100 characters
*   `custom_one` - 1-100 characters
*   `custom_two` - 1-100 characters
*   `custom_three` - 1-100 characters
*   `custom_four` - 1-100 characters
*   `custom_five` - 1-100 characters
*   `active` - `Y` or `N`
*   `wrapup_seconds_override` - from -1 to 9999
*   `campaign_rank` - from -9 to 9 (changes campaign rank for this user on all campaigns)
*   `campaign_grade` - from 1 to 10 (changes campaign grade on all campaigns)
*   `ingroup_rank` - from -9 to 9 (changes in-group rank on all in-groups)
*   `ingroup_grade` - from 1 to 10 (changes in-group grade on all in-groups)
*   `camp_rg_only` - `0` or `1` (updates only one campaign rank/grade, requires `campaign_id` to be set)
*   `campaign_id` - required if `camp_rg_only` is set
*   `ingrp_rg_only` - `0` or `1` (updates only one in-group rank/grade, requires `group_id` to be set)
*   `group_id` - required if `ingrp_rg_only` is set
*   `reset_password` - `0` or `8+` digits (resets password to random string of this length and emails it to the user. `email` field is required if user's current email is invalid)
*   `in_groups` - pipe-delimited list of inbound groups (e.g. `SALESLINE\|SUPPORT`). *Does not affect currently logged in sessions.*

#### Example URLs
```http
http://server/vicidial/non_agent_api.php?source=test&function=update_user&user=6666&pass=1234&agent_user=2424&agent_pass=9876
http://server/vicidial/non_agent_api.php?source=test&function=update_user&user=6666&pass=1234&agent_user=2424&agent_user_level=3
http://server/vicidial/non_agent_api.php?source=test&function=update_user&user=6666&pass=1234&agent_user=2424&agent_full_name=Testing&hotkeys_active=1&active=Y
http://server/vicidial/non_agent_api.php?source=test&function=update_user&user=6666&pass=1234&agent_user=2424&campaign_rank=3&campaign_grade=4
http://server/vicidial/non_agent_api.php?source=test&function=update_user&user=6666&pass=1234&agent_user=2424&ingroup_rank=3&ingroup_grade=4&ingrp_rg_only=1&group_id=TEST_IN
http://server/vicidial/non_agent_api.php?source=test&function=update_user&user=6666&pass=1234&agent_user=2424&reset_password=12&email=info@vicidial.com
```

#### Example Responses
```text
ERROR: update_user USER DOES NOT HAVE PERMISSION TO UPDATE USERS - 6666|0
ERROR: update_user USER DOES NOT HAVE PERMISSION TO DELETE USERS - 6666|7878
ERROR: update_user YOU MUST USE ALL REQUIRED FIELDS - 6666||
ERROR: update_user USER DOES NOT EXIST - 6666|78789
ERROR: update_user USER DOES NOT HAVE PERMISSION TO UPDATE THIS USER - 6666|7878
ERROR: update_user YOU MUST USE A VALID FULL NAME, THIS IS AN OPTIONAL FIELD - 6666|
ERROR: update_user YOU MUST USE A VALID USER LEVEL, THIS IS AN OPTIONAL FIELD - 6666|0|9|1
ERROR: update_user YOU MUST USE A VALID USER GROUP, THIS IS AN OPTIONAL FIELD - 6666|FAKEGROUP|0
ERROR: update_user YOU MUST USE A VALID PHONE LOGIN, THIS IS AN OPTIONAL FIELD - 6666|fakephone|0|0
ERROR: update_user YOU MUST USE A VALID PHONE PASSWORD, THIS IS AN OPTIONAL FIELD - 6666|0
ERROR: update_user YOU MUST USE A VALID HOTKEYS SETTING, THIS IS AN OPTIONAL FIELD - 6666|2
ERROR: update_user YOU MUST USE A VALID VOICEMAIL ID, THIS IS AN OPTIONAL FIELD - 6666
ERROR: update_user YOU MUST USE A VALID EMAIL, THIS IS AN OPTIONAL FIELD - 6666
ERROR: update_user YOU MUST USE A VALID CUSTOM ONE, THIS IS AN OPTIONAL FIELD - 6666
ERROR: update_user YOU MUST USE A VALID CUSTOM TWO, THIS IS AN OPTIONAL FIELD - 6666
ERROR: update_user YOU MUST USE A VALID CUSTOM THREE, THIS IS AN OPTIONAL FIELD - 6666
ERROR: update_user YOU MUST USE A VALID CUSTOM FOUR, THIS IS AN OPTIONAL FIELD - 6666
ERROR: update_user YOU MUST USE A VALID CUSTOM FIVE, THIS IS AN OPTIONAL FIELD - 6666
ERROR: update_user ACTIVE MUST BE Y OR N, THIS IS AN OPTIONAL FIELD - 6666
ERROR: update_user wrapup_seconds_override MUST BE A VALID DIGIT BETWEEN -1 AND 9999, THIS IS AN OPTIONAL FIELD - 6666|-2
ERROR: update_user YOU MUST USE A VALID CAMPAIGN RANK, THIS IS AN OPTIONAL FIELD - 6666|0
ERROR: update_user YOU MUST USE A VALID CAMPAIGN GRADE, THIS IS AN OPTIONAL FIELD - 6666|0
ERROR: update_user YOU MUST USE A VALID IN-GROUP RANK, THIS IS AN OPTIONAL FIELD - 6666|0
ERROR: update_user YOU MUST USE A VALID IN-GROUP GRADE, THIS IS AN OPTIONAL FIELD - 6666|0
ERROR: update_user NO UPDATES DEFINED - 6666|0
NOTICE: update_user USER CAMPAIGN RANKS/GRADES HAVE BEEN UPDATED - 6666|9|10|TESTCAMP|1
NOTICE: update_user USER IN-GROUP RANKS/GRADES HAVE BEEN UPDATED - 6666|9|10|TEST_IN|1
NOTICE: update_user USER PASSWORD RESET FAILED: NO VALID EMAIL PROVIDED - 6666|email@doman
NOTICE: update_user USER PASSWORD RESET FAILED: EMAIL FAILURE - 6666|info@vicidial.com|email-error
NOTICE: update_user update_user USER PASSWORD HAS BEEN RESET - 6666|info@vicidial.com
SUCCESS: update_user USER HAS BEEN UPDATED - 6666|7878
SUCCESS: update_user USER HAS BEEN DELETED - 6666|7878
```

---

### update_remote_agent
Updates remote agent settings in the system.

> [!NOTE]
> API user for this function must have `user_level` set to 8 or higher and "modify remote agents" enabled.

#### Required Fields
*   `agent_user` - 2-20 characters

#### Optional Fields
*   `status` - `ACTIVE` or `INACTIVE` only
*   `campaign_id` - 1-8 characters, must be a valid Campaign ID
*   `number_of_lines` - 1-20 digits

#### Example URL
```http
http://server/vicidial/non_agent_api.php?source=test&function=update_remote_agent&user=6666&pass=1234&agent_user=2424&number_of_lines=1&campaign_id=TESTCAMP
http://server/vicidial/non_agent_api.php?source=test&function=update_remote_agent&user=6666&pass=1234&agent_user=2424&status=INACTIVE
```

#### Example Responses
```text
ERROR: update_remote_agent USER DOES NOT HAVE PERMISSION TO UPDATE REMOTE AGENTS - 6666|0
ERROR: update_remote_agent YOU MUST USE ALL REQUIRED FIELDS - 6666||
ERROR: update_remote_agent REMOTE AGENT DOES NOT EXIST - 6666|78789
ERROR: update_remote_agent USER DOES NOT HAVE PERMISSION TO UPDATE THIS REMOTE AGENT - 6666
ERROR: update_remote_agent YOU MUST USE A VALID NUMBER OF LINES, THIS IS AN OPTIONAL FIELD - 6666|0
ERROR: update_remote_agent YOU MUST USE A VALID CAMPAIGN ID, THIS IS AN OPTIONAL FIELD - 6666|TEST|0
ERROR: update_remote_agent STATUS MUST BE ACTIVE OR INACTIVE, THIS IS AN OPTIONAL FIELD - 6666|Y
ERROR: update_remote_agent NO UPDATES DEFINED - 6666|0
NOTICE: update_remote_agent REMOTE AGENT IS ALREADY ACTIVE - 6666|7878
NOTICE: update_remote_agent REMOTE AGENT IS ALREADY INACTIVE - 6666|7878
SUCCESS: update_remote_agent USER HAS BEEN UPDATED - 6666|7878
```

---

### add_group_alias
Adds a group alias to the system.

> [!NOTE]
> API user for this function must have `user_level` set to 8 or higher and "ast admin access" enabled.

#### Required Fields
*   `caller_id_number` - 6-20 characters

#### Optional Fields
*   `group_alias_id` - 3-30 characters (no spaces or punctuation), defaults to `caller_id_number`
*   `group_alias_name` - 3-50 characters, defaults to `caller_id_number`
*   `caller_id_name` - 6-20 characters, defaults to empty `''`
*   `active` - `Y` or `N` (default `Y`)
*   `admin_user_group` - valid user group (default `---ALL---`)

#### Example URL
```http
http://server/vicidial/non_agent_api.php?source=test&function=add_group_alias&user=6666&pass=1234&caller_id_number=7275551212
http://server/vicidial/non_agent_api.php?source=test&function=add_group_alias&user=6666&pass=1234&caller_id_number=7275551212&group_alias_id=test_group_alias&group_alias_name=test+group+alias
```

#### Example Responses
```text
ERROR: add_group_alias USER DOES NOT HAVE PERMISSION TO ADD GROUP ALIASES - 6666|0
ERROR: add_group_alias YOU MUST USE ALL REQUIRED FIELDS - 6666|||1000
ERROR: add_group_alias GROUP ALIAS ALREADY EXISTS - 6666|7275551212
SUCCESS: add_group_alias GROUP ALIAS HAS BEEN ADDED - 6666|7275551212|test_group_alias|test group alias
```

---

### add_dnc_phone
Adds a phone number to a DNC list in the system.

> [!NOTE]
> API user for this function must have `user_level` set to 8 or higher and "modify lists" enabled.

#### Required Fields
*   `phone_number` - 6-20 characters
*   `campaign_id` - 2-30 characters, must be a valid campaign ID or `SYSTEM_INTERNAL`

#### Example URL
```http
http://server/vicidial/non_agent_api.php?source=test&function=add_dnc_phone&user=6666&pass=1234&phone_number=7275551212&campaign_id=TESTCAMP
http://server/vicidial/non_agent_api.php?source=test&function=add_dnc_phone&user=6666&pass=1234&phone_number=7275551212&campaign_id=SYSTEM_INTERNAL
```

#### Example Responses
```text
ERROR: add_dnc_phone USER DOES NOT HAVE PERMISSION TO ADD DNC NUMBERS - 6666|0
ERROR: add_dnc_phone YOU MUST USE ALL REQUIRED FIELDS - 6666|||1000
ERROR: add_dnc_phone DNC NUMBER ALREADY EXISTS - 6666|7275551212
SUCCESS: add_dnc_phone DNC NUMBER HAS BEEN ADDED - 6666|7275551212|TESTCAMP
```

---

### delete_dnc_phone
Removes a phone number from a DNC list in the system.

> [!NOTE]
> API user for this function must have `user_level` set to 8 or higher, "modify lists" enabled, and "Delete From DNC Lists" enabled.

#### Required Fields
*   `phone_number` - 6-20 characters
*   `campaign_id` - 2-30 characters, must be a valid campaign ID or `SYSTEM_INTERNAL`

#### Example URL
```http
http://server/vicidial/non_agent_api.php?source=test&function=delete_dnc_phone&user=6666&pass=1234&phone_number=7275551212&campaign_id=TESTCAMP
http://server/vicidial/non_agent_api.php?source=test&function=delete_dnc_phone&user=6666&pass=1234&phone_number=7275551212&campaign_id=SYSTEM_INTERNAL
```

#### Example Responses
```text
ERROR: delete_dnc_phone USER DOES NOT HAVE PERMISSION TO DELETE DNC NUMBERS - 6666|0
ERROR: delete_dnc_phone YOU MUST USE ALL REQUIRED FIELDS - 6666|||1000
ERROR: delete_dnc_phone DNC NUMBER DOES NOT EXIST - 6666|7275551212
SUCCESS: delete_dnc_phone DNC NUMBER HAS BEEN DELETED - 6666|7275551212|TESTCAMP
```

---

### add_fpg_phone
Adds a phone number to a Filter Phone Group in the system.

> [!NOTE]
> API user for this function must have `user_level` set to 8 or higher and "modify lists" enabled.

#### Required Fields
*   `phone_number` - 6-20 characters
*   `group` - 2-30 characters, must be a valid Filter Phone Group ID

#### Example URL
```http
http://server/vicidial/non_agent_api.php?source=test&function=add_fpg_phone&user=6666&pass=1234&phone_number=7275551212&group=TEST_IN_FILTER
```

#### Example Responses
```text
ERROR: add_fpg_phone USER DOES NOT HAVE PERMISSION TO ADD FILTER PHONE GROUP NUMBERS - 6666|0
ERROR: add_fpg_phone YOU MUST USE ALL REQUIRED FIELDS - 6666|||1000
ERROR: add_fpg_phone FILTER PHONE GROUP DOES NOT EXIST - 6666|TEST_IN_FILTER2
ERROR: add_fpg_phone FILTER PHONE GROUP NUMBER ALREADY EXISTS - 6666|7275551212|TEST_IN_FILTER
SUCCESS: add_fpg_phone FILTER PHONE GROUP NUMBER HAS BEEN ADDED - 6666|7275551212|TEST_IN_FILTER
```

---

### delete_fpg_phone
Removes a phone number from a Filter Phone Group in the system.

> [!NOTE]
> API user for this function must have `user_level` set to 8 or higher and "modify lists" enabled.

#### Required Fields
*   `phone_number` - 6-20 characters
*   `group` - 2-30 characters, must be a valid Filter Phone Group ID

#### Example URL
```http
http://server/vicidial/non_agent_api.php?source=test&function=delete_fpg_phone&user=6666&pass=1234&phone_number=7275551212&group=TEST_IN_FILTER
```

#### Example Responses
```text
ERROR: delete_fpg_phone USER DOES NOT HAVE PERMISSION TO DELETE FILTER PHONE GROUP NUMBERS - 6666|0
ERROR: delete_fpg_phone YOU MUST USE ALL REQUIRED FIELDS - 6666|||1000
ERROR: delete_fpg_phone FILTER PHONE GROUP DOES NOT EXIST - 6666|TEST_IN_FILTER2
ERROR: delete_fpg_phone FILTER PHONE GROUP NUMBER DOES NOT EXIST - 6666|7275551212|TEST_IN_FILTER
SUCCESS: delete_fpg_phone FILTER PHONE GROUP NUMBER HAS BEEN DELETED - 6666|7275551212|TEST_IN_FILTER
```

---

### add_phone
Adds a phone to the system.

> [!NOTE]
> API user for this function must have `user_level` set to 8 or higher and "ast admin access" enabled.

#### Required Fields
*   `extension` - 2-100 characters
*   `dialplan_number` - 1-20 digits
*   `voicemail_id` - 1-10 digits
*   `phone_login` - 1-20 characters
*   `phone_pass` - 1-20 characters
*   `server_ip` - must be a valid server_ip (7-15 characters)
*   `protocol` - `IAX2`, `SIP`, `Zap`, or `EXTERNAL`
*   `registration_password` - 1-20 characters
*   `phone_full_name` - 1-50 characters
*   `local_gmt` - timezone offset (no DST adjustment, default: `-5.00`)
*   `outbound_cid` - 1-20 digits

#### Optional Fields
*   `phone_context` - default is `'default'`
*   `email` - 1-100 characters
*   `admin_user_group` - valid user group (default `---ALL---`)
*   `is_webphone` - `Y`, `N` or `Y_API_LAUNCH` (default `N`)
*   `webphone_auto_answer` - `Y` or `N` (default `N`)
*   `use_external_server_ip` - `Y` or `N` (default `N`)
*   `template_id` - 1-15 characters (must be valid system template)
*   `on_hook_agent` - `Y` or `N` (default `N`)

#### Example URL
```http
http://server/vicidial/non_agent_api.php?source=test&function=add_phone&user=6666&pass=1234&extension=55000&dialplan_number=55000&voicemail_id=55000&phone_login=55000&phone_pass=test&server_ip=192.168.198.5&protocol=SIP&registration_password=test&phone_full_name=Extension+55000&local_gmt=-5.00&outbound_cid=7275551212
http://server/vicidial/non_agent_api.php?source=test&function=add_phone&user=6666&pass=1234&extension=55000&dialplan_number=55000&voicemail_id=55000&phone_login=55000&phone_pass=test&server_ip=192.168.198.5&protocol=SIP&registration_password=test&phone_full_name=Extension+55000&local_gmt=-5.00&outbound_cid=7275551212&phone_context=default&email=test@test.com
```

#### Example Responses
```text
ERROR: add_phone USER DOES NOT HAVE PERMISSION TO ADD PHONES - 6666|0
ERROR: add_phone YOU MUST USE ALL REQUIRED FIELDS - 6666|1000|||||||||||
ERROR: add_phone SERVER DOES NOT EXIST - 6666|10.0.9.9
ERROR: add_phone PHONE ALREADY EXISTS ON THIS SERVER - 6666|10.0.9.8|cc101
ERROR: add_phone PHONE LOGIN ALREADY EXISTS - 6666|cc100
ERROR: add_phone YOU MUST USE A VALID TIMEZONE - 6666|-5
ERROR: add_phone TEMPLATE ID DOES NOT EXIST, THIS IS AN OPTIONAL FIELD - 6666|cc100|batch
SUCCESS: add_phone PHONE HAS BEEN ADDED - 6666|cc100|10.0.9.8|SIP|100
```

---

### update_phone
Updates or deletes an existing phone entry in the system.

> [!NOTE]
> API user for this function must have `user_level` set to 8 or higher and "ast admin access" enabled.

#### Required Fields
*   `extension` - 2-100 characters
*   `server_ip` - must be a valid server_ip (7-15 characters)

#### Settings Fields
*   `delete_phone` - `Y` or `N` (default `N`). Deletes phone from the system if `Y`.

#### Editable Fields
*   `dialplan_number` - 1-20 digits
*   `voicemail_id` - 1-10 digits
*   `phone_login` - 1-20 characters
*   `phone_pass` - 1-20 characters
*   `protocol` - `IAX2`, `SIP`, `Zap`, or `EXTERNAL`
*   `registration_password` - 1-20 characters
*   `phone_full_name` - 1-50 characters
*   `local_gmt` - default: `-5.00`
*   `outbound_cid` - 6-18 digits
*   `outbound_alt_cid` - 6-18 digits (send `--BLANK--` to clear)
*   `phone_context` - default is `'default'`
*   `email` - 5-100 characters (send `--BLANK--` to clear)
*   `admin_user_group` - valid user group (default `---ALL---`)
*   `phone_ring_timeout` - 1-3 digits (default `'60'`)
*   `delete_vm_after_email` - `Y` or `N` (default `'N'`)
*   `is_webphone` - `Y`, `N` or `Y_API_LAUNCH` (default `N`)
*   `webphone_auto_answer` - `Y` or `N` (default `N`)
*   `use_external_server_ip` - `Y` or `N` (default `N`)
*   `template_id` - 1-15 characters (must be valid system template)
*   `on_hook_agent` - `Y` or `N` (default `N`)

#### Example URL
```http
http://server/vicidial/non_agent_api.php?source=test&function=update_phone&user=6666&pass=1234&extension=55000&dialplan_number=55000&voicemail_id=55000&phone_login=55000&phone_pass=test&server_ip=192.168.198.5&protocol=SIP&registration_password=test&phone_full_name=Extension+55000&local_gmt=-5.00&outbound_cid=7275551212
http://server/vicidial/non_agent_api.php?source=test&function=update_phone&user=6666&pass=1234&extension=55000&dialplan_number=55000&voicemail_id=55000&phone_login=55000&phone_pass=test&server_ip=192.168.198.5&protocol=SIP&registration_password=test&phone_full_name=Extension+55000&local_gmt=-5.00&outbound_cid=7275551212&phone_context=default&email=test@test.com
http://server/vicidial/non_agent_api.php?source=test&function=update_phone&user=6666&pass=1234&extension=55000&server_ip=192.168.198.5&delete_phone=Y
```

#### Example Responses
```text
ERROR: update_phone USER DOES NOT HAVE PERMISSION TO UPDATE PHONES - 6666|0
ERROR: update_phone YOU MUST USE ALL REQUIRED FIELDS - 6666||
ERROR: update_phone SERVER DOES NOT EXIST - 6666|10.0.9.9
ERROR: update_phone PHONE DOES NOT EXIST ON THIS SERVER - 6666|10.0.9.8|cc101
ERROR: update_phone LOGIN ALREADY EXISTS - 6666|x101
ERROR: update_phone YOU MUST USE A VALID TIMEZONE - 6666|-99
ERROR: update_phone YOU MUST USE A VALID DIALPLAN NUMBER - 6666|0
ERROR: update_phone YOU MUST USE A VALID VOICEMAIL NUMBER - 6666|0
ERROR: update_phone YOU MUST USE A VALID PHONE PASSWORD - 6666|0
ERROR: update_phone YOU MUST USE A VALID REGISTRATION PASSWORD - 6666|0
ERROR: update_phone YOU MUST USE A VALID PROTOCOL - 6666|0
ERROR: update_phone YOU MUST USE A VALID PHONE NAME - 6666|0
ERROR: update_phone YOU MUST USE A VALID PHONE CONTEXT - 6666|0
ERROR: update_phone YOU MUST USE A VALID EMAIL - 6666|x.com
ERROR: update_phone NO UPDATES DEFINED - 6666|0
ERROR: update_phone ACTIVE MUST BE Y OR N, THIS IS AN OPTIONAL FIELD - 6666|0
ERROR: update_phone OUTBOUND CID MUST BE FROM 6 TO 18 DIGITS, THIS IS AN OPTIONAL FIELD - 6666|0
ERROR: update_phone OUTBOUND ALT CID MUST BE FROM 6 TO 18 DIGITS, THIS IS AN OPTIONAL FIELD - 6666|0
ERROR: update_phone USER GROUP MUST BE FROM 2 TO 20 CHARACTERS, THIS IS AN OPTIONAL FIELD - 6666|0
ERROR: update_phone PHONE RING TIMEOUT MUST BE FROM 2 TO 180 SECONDS, THIS IS AN OPTIONAL FIELD - 6666|0
ERROR: update_phone DELETE VOICEMAIL AFTER EMAIL MUST BE Y OR N, THIS IS AN OPTIONAL FIELD - 6666|0
ERROR: update_phone TEMPLATE ID DOES NOT EXIST, THIS IS AN OPTIONAL FIELD - 6666|cc100|batch
SUCCESS: update_phone PHONE HAS BEEN UPDATED - 6666|cc100|10.0.9.8|SIP|100
SUCCESS: update_phone PHONE HAS BEEN DELETED - 6666|cc100|10.0.9.8|
NOTICE: update_phone USER DOES NOT HAVE PERMISSION TO DELETE PHONES - 6666|7878
```

---

### add_phone_alias
Adds a phone alias entry to the system.

> [!NOTE]
> API user for this function must have `user_level` set to 8 or higher and "ast admin access" enabled.

#### Required Fields
*   `alias_id` - 2-20 characters
*   `phone_logins` - 2-255 characters (phone logins separated by commas)
*   `alias_name` - 1-50 characters

#### Example URL
```http
http://server/vicidial/non_agent_api.php?source=test&function=add_phone_alias&user=6666&pass=1234&alias_id=xyz100&alias_name=XYZ+testing&phone_logins=100a,100b
```

#### Example Responses
```text
ERROR: add_phone_alias USER DOES NOT HAVE PERMISSION TO ADD PHONE ALIASES - 6666|0
ERROR: add_phone_alias YOU MUST USE ALL REQUIRED FIELDS - 6666||||
ERROR: add_phone_alias PHONE ALIAS ALREADY EXISTS - 6666|x101
ERROR: add_phone_alias PHONE DOES NOT EXIST - 6666|cc100
SUCCESS: add_phone_alias PHONE ALIAS HAS BEEN ADDED - 6666|x100|testing_x100|cc100,bb100
```

---

### update_phone_alias
Updates or deletes a phone alias entry already in the system.

> [!NOTE]
> API user for this function must have `user_level` set to 8 or higher and "ast admin access" enabled.

#### Required Fields
*   `alias_id` - 2-20 characters
*   `phone_logins` - 2-255 characters (separated by commas)
*   `alias_name` - 1-50 characters

#### Settings Fields
*   `delete_alias` - `Y` or `N` (default `N`). Deletes phone alias if set to `Y`.

#### Example URL
```http
http://server/vicidial/non_agent_api.php?source=test&function=update_phone_alias&user=6666&pass=1234&alias_id=xyz100&alias_name=XYZ+testing+X&phone_logins=100a,100b
http://server/vicidial/non_agent_api.php?source=test&function=update_phone_alias&user=6666&pass=1234&alias_id=xyz100&delete_alias=Y
```

#### Example Responses
```text
ERROR: update_phone_alias USER DOES NOT HAVE PERMISSION TO UPDATE PHONE ALIASES - 6666|0
ERROR: update_phone_alias YOU MUST USE ALL REQUIRED FIELDS - 6666||||
ERROR: update_phone_alias PHONE ALIAS DOES NOT EXIST - 6666|x101
ERROR: update_phone_alias YOU MUST USE A VALID ALIAS NAME - 6666|x
ERROR: update_phone_alias PHONE DOES NOT EXIST - 6666|cc100
SUCCESS: update_phone_alias PHONE ALIAS HAS BEEN UPDATED - 6666|x100|
SUCCESS: update_phone_alias PHONE ALIAS HAS BEEN DELETED - 6666|x100|
```

---

### server_refresh
Forces a conf file refresh on all telco servers in the cluster.

> [!NOTE]
> API user for this function must have `user_level` set to 8 or higher and "ast admin access" enabled.

#### Required Fields
*   `stage` - `"REFRESH"`

#### Example URL
```http
http://server/vicidial/non_agent_api.php?source=test&function=server_refresh&user=6666&pass=1234&stage=REFRESH
```

#### Example Responses
```text
ERROR: server_refresh USER DOES NOT HAVE PERMISSION TO REFRESH SERVERS - 6666|0
ERROR: server_refresh YOU MUST USE ALL REQUIRED FIELDS - 6666|
SUCCESS: server_refresh SERVER REFRESH HAS BEEN TRIGGERED - 6666|1
```

---

### add_list
Adds a list to the system.

> [!NOTE]
> API user for this function must have `user_level` set to 8 or higher and "modify lists" enabled.
> * Web form addresses can be URL encoded.
> * To use `custom_fields_copy`, the API user must have the "Custom Fields Modify" user option enabled.

#### Required Fields
*   `list_id` - 2-14 digits
*   `list_name` - 2-30 characters
*   `campaign_id` - 2-8 characters (must be a valid campaign)

#### Optional Fields
*   `active` - `Y` or `N` (default `N`)
*   `list_description` - up to 255 characters (use `--BLANK--` to empty)
*   `outbound_cid` - 6-20 digits
*   `script` - 1-10 characters (must be valid system script)
*   `am_message` - 2-100 characters
*   `drop_inbound_group` - 1-10 characters (must be valid in-group)
*   `web_form_address` - 6-100 characters
*   `web_form_address_two` - 6-100 characters
*   `web_form_address_three` - 6-100 characters
*   `reset_time` - 4-100 characters, valid 4-digit groups of 24-hour time (e.g. `0900-1700-2359`)
*   `tz_method` - `COUNTRY_AND_AREA_CODE`, `POSTAL_CODE`, `NANPA_PREFIX`, or `OWNER_TIME_ZONE_CODE` (default `COUNTRY_AND_AREA_CODE`)
*   `local_call_time` - must be a valid call time ID, or `'campaign'` (default)
*   `expiration_date` - 10 characters, YYYY-MM-DD format (e.g. `2012-11-25`)
*   `xferconf_one` - `xferconf_five` - Transfer-Conf Number Overrides: 1-50 characters
*   `custom_fields_copy` - A valid list ID to copy custom fields from
*   `custom_copy_method` - `APPEND`, `UPDATE`, or `REPLACE` (default `APPEND`)
    *   `APPEND` - only adds fields not present in the new list (safest method)
    *   `UPDATE` - updates only existing custom field definitions to match source
    *   `REPLACE` - deletes all existing custom field lead data in new list and starts clean

#### Example URLs
```http
http://server/vicidial/non_agent_api.php?source=test&function=add_list&user=6666&pass=1234&list_id=1101&list_name=Test+API+list&campaign_id=TESTCAMP
http://server/vicidial/non_agent_api.php?source=test&function=add_list&user=6666&pass=1234&list_id=1101&list_name=Test+API+list&campaign_id=TESTCAMP&active=N&outbound_cid=7275551212&script=DEMOSCRIPT&am_message=8304&drop_inbound_group=SALESLINE&web_form_address=http://www.vicidial.org/?testing=hghg
http://server/vicidial/non_agent_api.php?source=test&user=6666&pass=1234&function=add_list&list_id=303&list_name=Test+API+list+3&campaign_id=TESTCAMP&custom_fields_copy=222
```

#### Example Responses
```text
ERROR: add_list USER DOES NOT HAVE PERMISSION TO ADD LISTS - 6666|0
ERROR: add_list YOU MUST USE ALL REQUIRED FIELDS - 6666|1000||
ERROR: add_list CAMPAIGN DOES NOT EXIST - 6666|TESTCIMP
ERROR: add_list LIST ALREADY EXISTS - 6666|1101
ERROR: add_list SCRIPT DOES NOT EXIST, THIS IS AN OPTIONAL FIELD - 6666|TESTSCRIPT
ERROR: add_list IN-GROUP DOES NOT EXIST, THIS IS AN OPTIONAL FIELD - 6666|TEST_IN8
ERROR: add_list RESET TIME IS NOT VALID, THIS IS AN OPTIONAL FIELD - 6666|012
ERROR: add_list EXPIRATION DATE IS NOT VALID, THIS IS AN OPTIONAL FIELD - 6666|012
ERROR: add_list LOCAL CALL TIME DOES NOT EXIST, THIS IS AN OPTIONAL FIELD - 6666|012
ERROR: add_list TIME ZONE METHOD IS NOT VALID, THIS IS AN OPTIONAL FIELD - 6666|012
ERROR: add_list TRANSFER CONF OVERRIDE ONE IS NOT VALID, THIS IS AN OPTIONAL FIELD - 6666|123456789012345678901234567890123456789012345678901
ERROR: add_list TRANSFER CONF OVERRIDE TWO IS NOT VALID, THIS IS AN OPTIONAL FIELD - 6666|123456789012345678901234567890123456789012345678901
ERROR: add_list TRANSFER CONF OVERRIDE THREE IS NOT VALID, THIS IS AN OPTIONAL FIELD - 6666|123456789012345678901234567890123456789012345678901
ERROR: add_list TRANSFER CONF OVERRIDE FOUR IS NOT VALID, THIS IS AN OPTIONAL FIELD - 6666|123456789012345678901234567890123456789012345678901
ERROR: add_list TRANSFER CONF OVERRIDE FIVE IS NOT VALID, THIS IS AN OPTIONAL FIELD - 6666|123456789012345678901234567890123456789012345678901
ERROR: add_list CUSTOM FIELDS LIST ID TO COPY FROM DOES NOT EXIST, THIS IS AN OPTIONAL FIELD - 6666|1101|1102|0
ERROR: add_list CUSTOM FIELDS LIST ID TO COPY FROM HAS NO CUSTOM FIELDS, THIS IS AN OPTIONAL FIELD - 6666|1101|1102|0
ERROR: add_list USER DOES NOT HAVE PERMISSION TO MODIFY CUSTOM FIELDS, THIS IS AN OPTIONAL FIELD - 6666|1101|1102|0
SUCCESS: add_list LIST HAS BEEN ADDED - 6666|1101|TESTCAMP
```

---

### update_list
Updates list information in the `vicidial_lists` table.

> [!NOTE]
> API user for this function must have `user_level` set to 8 or higher and "modify lists" enabled.
> * To set a field to empty, use the value `--BLANK--` (e.g. `&drop_inbound_group=--BLANK--`).
> * Avoid special characters like apostrophes, quotes, or ampersands.

#### Required Fields
*   `list_id` - must be all numbers, 2-14 digits
*   `source` - description of what originated the API call (maximum 20 characters)

#### Settings Fields
*   `insert_if_not_found` - `Y` or `N` (default `N`). Will attempt to insert as a new list if no match is found (switches to `add_list` processing).
*   `reset_list` - `Y` or `N` (default `N`). Resets called flags for all leads in the list.
*   `delete_list` - `Y` or `N` (default `N`). Deletes list from `vicidial_lists`.
*   `delete_leads` - `Y` or `N` (default `N`). Deletes all leads with this `list_id`.
*   `archived_lead` - `Y` or `N` (default `N`). Affects only archived tables if `Y`.

#### Editable Fields
*   `list_name` - 6-30 characters
*   `list_description` - up to 255 characters (use `--BLANK--` to empty)
*   `campaign_id` - 2-8 characters
*   `active` - `Y` or `N` (default `N`)
*   `outbound_cid` - 6-20 digits
*   `script` - 1-10 characters (must be valid system script)
*   `am_message` - 2-100 characters
*   `drop_inbound_group` - 1-10 characters (must be valid in-group)
*   `web_form_address` - 6-100 characters
*   `web_form_address_two` - 6-100 characters
*   `web_form_address_three` - 6-100 characters
*   `reset_time` - 4-100 characters (e.g. `0900-1700-2359`)
*   `tz_method` - `COUNTRY_AND_AREA_CODE`, `POSTAL_CODE`, `NANPA_PREFIX`, or `OWNER_TIME_ZONE_CODE`
*   `local_call_time` - system call time ID or `'campaign'`
*   `expiration_date` - YYYY-MM-DD
*   `xferconf_one` - `xferconf_five` - Transfer-Conf Overrides: 1-50 characters
*   `custom_fields_copy` - valid List ID with custom fields to copy
*   `custom_copy_method` - `APPEND`, `UPDATE`, or `REPLACE`
*   `custom_fields_add` - `Y` or `N` (default `N`). Requires fields marked below with an asterisk (\*):
    *   `field_label`\* - label used in database (letters, numbers, underscores only)
    *   `field_name`\* - visible name of field in form
    *   `field_size`\* - size of form field
    *   `field_type`\* - type of field (e.g. `TEXT`, `SELECT`, `AREA`, etc.)
    *   `field_rank`\* - top-to-bottom order in form
    *   `field_order` - horizontal order of field (if rank is identical)
    *   `field_rerank` - move existing fields of same rank down: `YES` or `NO` (default `NO`)
    *   `field_max` - maximum size of data
    *   `field_default` - default value
    *   `field_options` - options list. For new lines, use `%0A`; for pipes, use `%7C`
    *   `field_duplicate` - duplicate of another field: `Y` or `N` (default `N`)
    *   `field_description` - description of field
    *   `field_help` - agent help text
    *   `field_required` - field is required: `Y` or `N` (default `N`)
    *   `multi_position` - element position: `HORIZONTAL` or `VERTICAL` (default `HORIZONTAL`)
    *   `name_position` - name label position: `TOP` or `LEFT` (default `LEFT`)
    *   `field_encrypt` - encrypt this field (VICIhost only): `Y` or `N` (default `N`)
    *   `field_show_hide` - hidden fields option (VICIhost only): `DISABLED`, `X_OUT_ALL`, `LAST_1`, `LAST_2`, `LAST_3`, `LAST_4`, `FIRST_1_LAST_4`
*   `custom_fields_update` - `Y` or `N` (default `N`). Requires `field_label`.
    *   *Can set the following fields to `--BLANK--` to change to an empty value:* `field_description`, `field_help`, `field_options`, `field_default`
*   `custom_fields_delete` - `Y` or `N` (default `N`). Requires `field_label`.

#### Example URLs
```http
http://server/vicidial/non_agent_api.php?source=test&user=6666&pass=1234&function=update_list&list_id=1101&list_name=Updated+test+API+list
http://server/vicidial/non_agent_api.php?source=test&user=6666&pass=1234&function=update_list&list_id=1101&active=N
http://server/vicidial/non_agent_api.php?source=test&user=6666&pass=1234&function=update_list&list_id=1102&active=N&insert_if_not_found=Y&list_name=Updated+test+API+list&campaign_id=TESTCAMP&outbound_cid=7275551212&script=DEMOSCRIPT&am_message=8304&drop_inbound_group=SALESLINE&web_form_address=http://www.vicidial.org/?testing=hghg
http://server/vicidial/non_agent_api.php?source=test&user=6666&pass=1234&function=update_list&list_id=1102&reset_list=Y
http://server/vicidial/non_agent_api.php?source=test&user=6666&pass=1234&function=update_list&list_id=1102&delete_list=Y&delete_leads=Y
http://server/vicidial/non_agent_api.php?source=test&user=6666&pass=1234&function=update_list&list_id=1102&campaign_id=ASTRICON
http://server/vicidial/non_agent_api.php?source=test&user=6666&pass=1234&function=update_list&list_id=99883&custom_copy_method=APPEND&custom_fields_copy=9988
http://server/vicidial/non_agent_api.php?source=test&user=6666&pass=1234&function=update_list&list_id=9988&custom_fields_add=Y&field_label=new_api_field4&field_name=new+API+field4&field_size=20&field_type=TEXT&field_rank=16
http://server/vicidial/non_agent_api.php?source=test&user=6666&pass=1234&function=update_list&list_id=9988&custom_fields_update=Y&field_label=Utility&field_options=1,1_1_hello%0A2,2_2_hello%0A3,3_3_hello
http://server/vicidial/non_agent_api.php?source=test&user=6666&pass=1234&function=update_list&list_id=9988&custom_fields_update=Y&field_label=bill_period3&field_options=source=%3Ebill_start_month3%0Avalue=%3ESeptember%0Aoption=%3E%7Cselect%20bill-period-1%0Aoption=%3E1_month%7C1%20Month%0Aoption=%3E2_months%7C2%20Months%0Aoption=%3E3_months%7C3%20Months%0Avalue=%3E%0Aoption=%3E%7Cno%20match
http://server/vicidial/non_agent_api.php?source=test&user=6666&pass=1234&function=update_list&list_id=5040379&custom_fields_delete=Y&field_label=new_api_field6
```

#### Example Responses
```text
ERROR: update_list USER DOES NOT HAVE PERMISSION TO UPDATE LISTS - 6666
ERROR: update_list YOU MUST USE ALL REQUIRED FIELDS - 6666|1
ERROR: update_list NOT AN ALLOWED LIST ID - 98762
NOTICE: update_list LIST DOES NOT EXIST, SENDING TO add_list FUNCTION - 6666|1000
ERROR: update_list LIST DOES NOT EXIST - 6666|1000
ERROR: update_list CAMPAIGN DOES NOT EXIST, THIS IS AN OPTIONAL FIELD - 6666|TESTCIMP
ERROR: update_list SCRIPT DOES NOT EXIST, THIS IS AN OPTIONAL FIELD - 6666|TESTSCRIPT
ERROR: update_list IN-GROUP DOES NOT EXIST, THIS IS AN OPTIONAL FIELD - 6666|TEST_IN8
ERROR: update_list LIST NAME MUST BE FROM 6 TO 30 CHARACTERS, THIS IS AN OPTIONAL FIELD - 6666|test
ERROR: update_list ACTIVE MUST BE Y OR N, THIS IS AN OPTIONAL FIELD - 6666|U
ERROR: update_list OUTBOUND CID MUST BE FROM 6 TO 18 DIGITS, THIS IS AN OPTIONAL FIELD - 6666|12345
ERROR: update_list ANSWERING MACHINE MESSAGE MUST LESS THAN 101 DIGITS, THIS IS AN OPTIONAL FIELD - 6666|123...
ERROR: update_list RESET TIME IS NOT VALID, THIS IS AN OPTIONAL FIELD - 6666|012
ERROR: update_list EXPIRATION DATE IS NOT VALID, THIS IS AN OPTIONAL FIELD - 6666|012
ERROR: update_list TIME ZONE METHOD IS NOT VALID, THIS IS AN OPTIONAL FIELD - 6666|012
ERROR: update_list LOCAL CALL TIME DOES NOT EXIST, THIS IS AN OPTIONAL FIELD - 6666|012
ERROR: update_list TRANSFER CONF OVERRIDE ONE IS NOT VALID, THIS IS AN OPTIONAL FIELD - 6666|123456789012345678901234567890123456789012345678901
ERROR: update_list TRANSFER CONF OVERRIDE TWO IS NOT VALID, THIS IS AN OPTIONAL FIELD - 6666|123456789012345678901234567890123456789012345678901
ERROR: update_list TRANSFER CONF OVERRIDE THREE IS NOT VALID, THIS IS AN OPTIONAL FIELD - 6666|123456789012345678901234567890123456789012345678901
ERROR: update_list TRANSFER CONF OVERRIDE FOUR IS NOT VALID, THIS IS AN OPTIONAL FIELD - 6666|123456789012345678901234567890123456789012345678901
ERROR: update_list TRANSFER CONF OVERRIDE FIVE IS NOT VALID, THIS IS AN OPTIONAL FIELD - 6666|123456789012345678901234567890123456789012345678901
NOTICE: update_list NO UPDATES DEFINED - 6666|
SUCCESS: update_list LIST HAS BEEN UPDATED - 6666|1101
NOTICE: update_list LEADS IN LIST HAVE BEEN RESET - 6666|1101|324
NOTICE: update_list LIST RESET FAILED, daily reset limit reached: 2 / 2 - 6666|1101
NOTICE: update_list USER DOES NOT HAVE PERMISSION TO DELETE LISTS - 6666|0
NOTICE: update_list LIST HAS BEEN DELETED - 6666|1101|1
NOTICE: update_list USER DOES NOT HAVE PERMISSION TO DELETE LEADS - 6666|0
NOTICE: update_list LEADS IN LIST HAVE BEEN DELETED - 6666|1101|324
NOTICE: update_list CUSTOM FIELDS LIST ID TO COPY FROM DOES NOT EXIST, THIS IS AN OPTIONAL FIELD - 6666|1101|1102|0
NOTICE: update_list CUSTOM FIELDS LIST ID TO COPY FROM HAS NO CUSTOM FIELDS, THIS IS AN OPTIONAL FIELD - 6666|1101|1102|0
NOTICE: update_list USER DOES NOT HAVE PERMISSION TO MODIFY CUSTOM FIELDS, THIS IS AN OPTIONAL FIELD - 6666|1101|1102|0
NOTICE: update_list COPY CUSTOM FIELDS COMMAND SENT - 6666|1101|APPEND|ERROR: Fields not copied|
NOTICE: update_list COPY CUSTOM FIELDS COMMAND SENT - 6666|1101|APPEND|SUCCESS: Fields copied|
NOTICE: update_list CUSTOM FIELDS LIST ID TO ADD TO HAS NO CUSTOM FIELDS, THIS IS AN OPTIONAL FIELD - 6666|1101
NOTICE: update_list REQUIRED CUSTOM FIELDS VARIABLES ARE MISSING, FIELD NOT ADDED, THIS IS AN OPTIONAL FIELD - 6666|1101||||||
NOTICE: update_list CUSTOM FIELDS LIST ID TO UPDATE TO HAS NO CUSTOM FIELDS, THIS IS AN OPTIONAL FIELD - 6666|1101||||||
NOTICE: update_list REQUIRED CUSTOM FIELDS VARIABLES ARE MISSING, FIELD NOT UPDATED, THIS IS AN OPTIONAL FIELD - 6666|1101|||
NOTICE: update_list FIELD DOES NOT EXIST, FIELD NOT UPDATED, THIS IS AN OPTIONAL FIELD - 6666|1101|abc|||||
NOTICE: update_list CUSTOM FIELDS LIST ID TO DELETE TO HAS NO CUSTOM FIELDS, THIS IS AN OPTIONAL FIELD - 6666|1101|
NOTICE: update_list REQUIRED CUSTOM FIELDS VARIABLES ARE MISSING, FIELD NOT DELETED, THIS IS AN OPTIONAL FIELD - 6666|1101|||
NOTICE: update_list FIELD DOES NOT EXIST, FIELD NOT DELETED, THIS IS AN OPTIONAL FIELD - 6666|1101|abc|
NOTICE: update_list ADD CUSTOM FIELD COMMAND SENT - 6666|1101|new_field|TEXT|ERROR: Field not added|
NOTICE: update_list ADD CUSTOM FIELD COMMAND SENT - 6666|1101|new_field|TEXT|SUCCESS: Field added|
NOTICE: update_list UPDATE CUSTOM FIELD COMMAND SENT - 6666|1101|new_field|TEXT|ERROR: Field not updated|
NOTICE: update_list UPDATE CUSTOM FIELD COMMAND SENT - 6666|1101|new_field|TEXT|SUCCESS: Field updated|
NOTICE: update_list DELETE CUSTOM FIELD COMMAND SENT - 6666|1101|new_field|TEXT|ERROR: Field not deleted|
NOTICE: update_list DELETE CUSTOM FIELD COMMAND SENT - 6666|1101|new_field|TEXT|SUCCESS: Field deleted|
```

---

### list_info
Summary information about a list.

> [!NOTE]
> API user for this function must have `user_level` set to 8 or higher and "modify lists" enabled.

#### Required Fields
*   `list_id` - must be all numbers, 2-14 digits
*   `source` - description of what originated the API call (maximum 20 characters)

#### Optional Settings Fields
*   `leads_counts` - `Y` or `N` (default `N`). Includes count of total leads and NEW leads.
*   `dialable_count` - `Y` or `N` (default `N`). Provides counts of dialable leads based on campaign settings.
*   `stage` - format of exported data: `csv`, `tab`, `pipe` (default)
*   `header` - include a header (`YES`) or not (`NO`). Optional, default is `NO`.
*   `archived_lead` - `Y` or `N` (default `N`). Query archived tables if `Y`.

#### Example URL
```http
http://server/vicidial/non_agent_api.php?source=test&user=6666&pass=1234&function=list_info&list_id=1101&leads_counts=Y&header=YES
```

#### Example Responses
```text
ERROR: list_info USER DOES NOT HAVE PERMISSION TO MODIFY LISTS - 6666
ERROR: list_info YOU MUST USE ALL REQUIRED FIELDS - 6666|1
ERROR: list_info NOT AN ALLOWED LIST ID - 98762
ERROR: list_info LIST DOES NOT EXIST - 98762
```

##### Successful Output Format (Header Enabled)
| list_id | list_name | campaign_id | active | list_changedate | list_lastcalldate | expiration_date | resets_today |
| :--- | :--- | :--- | :--- | :--- | :--- | :--- | :--- |
| 1101 | performance list | TESTCAMP | Y | 2019-09-02 08:52:50 | 2019-11-07 10:12:12 | 2099-12-31 | 0 |

##### Successful Output Format (with `leads_counts=Y`)
| list_id | list_name | campaign_id | active | list_changedate | list_lastcalldate | expiration_date | resets_today | all_leads_count | new_leads_count |
| :--- | :--- | :--- | :--- | :--- | :--- | :--- | :--- | :--- | :--- |
| 1101 | performance list | TESTCAMP | Y | 2019-09-02 08:52:50 | 2019-11-07 10:12:12 | 2099-12-31 | 0 | 59896 | 3 |

---

### list_custom_fields
Shows custom data fields in a list, or across all lists.

> [!NOTE]
> API user for this function must have `user_level` set to 8 or higher and "modify lists" enabled.
> This output does not include display-only elements; only non-default custom fields are returned.

#### Required Fields
*   `list_id` - must be all numbers (2-14 digits) OR `---ALL---` for all lists
*   `source` - description of what originated the API call (maximum 20 characters)

#### Optional Settings Fields
*   `stage` - format of exported data: `csv`, `tab`, `pipe` (default)
*   `header` - include header (`YES`) or not (`NO`). Default is `NO`.
*   `custom_order` - order of custom fields: `table_order`, `alpha_up`, or `alpha_down` (default `'table_order'`)

#### Example URL
```http
http://server/vicidial/non_agent_api.php?source=test&user=6666&pass=1234&function=list_custom_fields&list_id=1101&header=YES
```

#### Example Responses
```text
ERROR: list_custom_fields CUSTOM LIST FIELDS ARE NOT ENABLED ON THIS SYSTEM - 6666
ERROR: list_custom_fields USER DOES NOT HAVE PERMISSION TO MODIFY LISTS - 6666
ERROR: list_custom_fields YOU MUST USE ALL REQUIRED FIELDS - 6666|1
ERROR: list_custom_fields NOT AN ALLOWED LIST ID - 98762
ERROR: list_custom_fields LIST DOES NOT EXIST - 98762
```

##### Successful Output Format (Header Enabled)
| list_id | list_name | campaign_id | active | list_changedate | list_lastcalldate | custom_fields_list_space_delimited |
| :--- | :--- | :--- | :--- | :--- | :--- | :--- |
| 1101 | performance list | TESTCAMP | Y | 2019-09-02 08:52:50 | 2019-11-07 10:12:12 | fav_color fav_food res_city res_state res_zip |

---

### check_phone_number
Allows checking if a phone number is valid and dialable.

> [!NOTE]
> API user for this function must have `user_level` set to 8 or higher.

#### Required Fields
*   `phone_number` - must be all numbers, 6-16 digits
*   `phone_code` - must be all numbers, 1-4 digits (defaults to 1 if not set)
*   `local_call_time` - must be a valid call time ID in the system

#### Optional Fields
*   `postal_code` - must be 5 digits (only works for US zip codes, required for `POSTAL` `tz_method`)
*   `state` - 2 letters (required for state call time rules)
*   `owner` - 2-5 letters (required for `TZCODE` `tz_method`)

#### Settings Fields
*   `dnc_check` - `Y`, `N` or `AREACODE` (default `N`)
*   `campaign_dnc_check` - `Y`, `N` or `AREACODE` (default `N`)
*   `campaign_id` - required if `campaign_dnc_check` is enabled
*   `usacan_prefix_check` - `Y` or `N` (default `N`). Check for valid US/Canada 4th digit.
*   `usacan_areacode_check` - `Y` or `N` (default `N`). Check for valid US/Canada areacode.
*   `nanpa_ac_prefix_check` - `Y` or `N` (default `N`). Check valid NANPA areacode and prefix.
*   `tz_method` - `<empty>`, `POSTAL`, `TZCODE`, `NANPA` or `AREACODE` (default `<empty>`, uses country code/areacode. Can combine methods, e.g. `POSTALTZCODE`)

#### Example URL
```http
http://server/vicidial/non_agent_api.php?source=test&user=6666&pass=1234&function=check_phone_number&phone_number=7275551111&call_time=9am-9pm

http://server/vicidial/non_agent_api.php?DB=0&source=test&function=check_phone_number&user=6666&pass=1234&owner=GST&postal_code=81009&phone_number=3125551212&local_call_time=9am-9pm&tz_method=AREACODEPOSTALNANPATZCODE
```

#### Example Responses
```text
ERROR: check_phone_number USER DOES NOT HAVE PERMISSION TO CHECK PHONE NUMBERS - |6666|0|
ERROR: check_phone_number THIS IS NOT A VALID CALL TIME - |6666|badct|
ERROR: check_phone_number NANPA options disabled, NANPA prefix data not loaded - 0|6666
ERROR: check_phone_number INVALID PHONE NUMBER LENGTH - 72755|6666 
ERROR: check_phone_number INVALID PHONE NUMBER PREFIX - 72755|6666 
ERROR: check_phone_number INVALID PHONE NUMBER AREACODE - 72755|6666 
ERROR: check_phone_number INVALID PHONE NUMBER NANPA AREACODE PREFIX - 7275551212|6666
ERROR: check_phone_number PHONE NUMBER IN DNC - 7275551112|6666
```

##### Successful Output Format
*(Returns results for each `tz_method` used as `[dialable(0/1)]|[gmt-offset]` followed by `#` and lead info)*

```text
AREACODE: 1|-6#PHONE: 3125551212|1||||
AREACODE: 1|-6#POSTAL: 1|-7.0#TZCODE: 0|10#NANPA: 1|-6#PHONE: 3125551212|1|81009||GST|
```

---

### logged_in_agents
List of agents currently logged in to the system.

> [!NOTE]
> API user for this function must have `user_level` set to 7 or higher and "view reports" enabled.

#### Required Fields
*   `source` - description of what originated the API call (maximum 20 characters)

#### Optional Fields
*   `campaigns` - pipe-delimited list of campaigns (e.g. `TESTCAMP\|INBOUND`, default is all campaigns)
*   `user_groups` - pipe-delimited list of user groups (e.g. `ADMIN\|AGENTS`, default is all groups)
*   `show_sub_status` - `YES` or `NO` (default `NO`). Show agent sub-status and `pause_code` (requires log lookup).

#### Settings Fields
*   `stage` - format of exported data: `csv`, `tab`, `pipe` (default)
*   `header` - include header (`YES`) or not (`NO`). Default is `NO`.

#### Example URL
```http
http://server/vicidial/non_agent_api.php?source=test&user=6666&pass=1234&function=logged_in_agents&stage=csv&header=YES
```

#### Example Responses
```text
ERROR: logged_in_agents USER DOES NOT HAVE PERMISSION TO GET AGENT INFO - 6666|0
ERROR: logged_in_agents NO LOGGED IN AGENTS
```

##### Successful Output (CSV format, default fields)
| user | campaign_id | session_id | status | lead_id | callerid | calls_today | full_name | user_group | user_level |
| :--- | :--- | :--- | :--- | :--- | :--- | :--- | :--- | :--- | :--- |
| 6666 | TESTCAMP | 8600051 | PAUSED | 1079409 | | 1 | Admin | ADMIN | 9 |

##### Successful Output (CSV format, sub-status enabled)
| user | campaign_id | session_id | status | lead_id | callerid | calls_today | full_name | user_group | user_level | pause_code | sub_status |
| :--- | :--- | :--- | :--- | :--- | :--- | :--- | :--- | :--- | :--- | :--- | :--- |
| 6666 | TESTCAMP | 8600051 | INCALL | 1079409 | M2260919190001079409 | 1 | Admin | ADMIN | 9 | LOGIN | DEAD |
| 4545 | TESTCAMP | 8600052 | PAUSED | 0 | | 0 | 4545 | MIKESGROUP | 8 | LOGIN | |

> [!NOTE]
> The `sub_status` field can consist of: `DEAD`, `DISPO`, `3-WAY`, `PARK`, `RING`, `PREVIEW`, `DIAL` or empty.

---

### call_status_stats
Report on the number of calls made by campaign and inbound group with hourly and status breakdowns.

> [!NOTE]
> API user for this function must have `user_level` set to 8 or higher and "view reports" enabled.

#### Required Fields
*   `campaigns` - Campaigns to return stats on. Use `---ALL---` or `ALLCAMPAIGNS` for all, or use dash delimiters for multiple campaigns.

#### Optional Fields
*   `query_date` - date to report on (YYYY-MM-DD, defaults to today's date)
*   `ingroups` - list of inbound groups, dash-delimited (defaults to all groups belonging to selected campaigns)
*   `statuses` - list of specific statuses to report on, dash-delimited (defaults to all)

#### Example URLs
```http
http://server/vicidial/non_agent_api.php?user=6666&pass=1234&function=call_status_stats&campaigns=---ALL---&query_date=2017-11-24
http://server/vicidial/non_agent_api.php?user=6666&pass=1234&function=call_status_stats&campaigns=TESTCAMP-TESTCAMP2&query_date=2017-11-24&statuses=NP-TD&ingroups=AGENTDIRECT
```

#### Example Responses
```text
ERROR: call_status_stats INVALID OR MISSING CAMPAIGNS
ERROR: auth USER DOES NOT HAVE PERMISSION TO USE THIS FUNCTION
ERROR: call_status_stats USER DOES NOT HAVE PERMISSION TO VIEW STATS
```

##### Successful Output Format (Pipe Delimited)
| campaign_id/ingroup | total calls | human answered calls | hourly breakdown (HH-count) | status breakdown (status-count) |
| :--- | :--- | :--- | :--- | :--- |
| TESTCAMP | 26 | 26 | 00-0,01-1,02-1,03-0,04-0,05-0,06-0,07-0,08-0,09-0,10-0,11-3,12-4,13-2,14-4,15-1,16-0,17-2,18-0,19-0,20-0,21-6,22-1,23-1 | NP-4,TD-22 |
| AGENTDIRECT | 2 | 2 | 00-0,01-0,02-0,03-0,04-0,05-0,06-0,07-0,08-0,09-0,10-0,11-1,12-1,13-0,14-0,15-0,16-0,17-0,18-0,19-0,20-0,21-0,22-0,23-0 | TD-2 |

---

### call_dispo_report
Call disposition breakdown report.

> [!NOTE]
> API user for this function must have `user_level` set to 9 and "view reports" enabled.

#### Required Fields (At least one of the following)
*   `campaigns` - Campaigns to return stats on, separated by dashes
*   `ingroups` - In-Groups to return stats on, separated by dashes
*   `dids` - DIDs to return stats on, separated by dashes

#### Optional Fields
*   `query_date` - start date of report (YYYY-MM-DD, default is today)
*   `end_date` - end date of report (YYYY-MM-DD, default is today)
*   `statuses` - list of specific statuses, dash-delimited
*   `categories` - list of status categories, dash-delimited
*   `users` - list of specific users, dash-delimited
*   `status_breakdown` - `0` or `1` (default `0`). Breakdown of all statuses.
*   `show_percentages` - `0` or `1` (default `0`). Shows percentages of statuses (requires `status_breakdown`).
*   `file_download` - `0` or `1` (default `0`). Download as a CSV file.

#### Example URLs
```http
http://server/vicidial/non_agent_api.php?source=test&user=6666&pass=1234&function=call_dispo_report&ingroups=TEST_IN3-TEST_IN2&query_date=2018-02-05&end_date=2018-02-07
http://server/vicidial/non_agent_api.php?source=test&user=6666&pass=1234&function=call_dispo_report&ingroups=TEST_IN3-TEST_IN2&query_date=2018-02-05&end_date=2018-02-07&status_breakdown=1
http://server/vicidial/non_agent_api.php?source=test&user=6666&pass=1234&function=call_dispo_report&did_patterns=7275551113&query_date=2018-02-05&end_date=2018-02-07&status_breakdown=1
http://server/vicidial/non_agent_api.php?source=test&user=6666&pass=1234&function=call_dispo_report&ingroups=TEST_IN3-TEST_IN2&query_date=2018-02-05&end_date=2018-02-07&status_breakdown=1&show_percentages=1
```

#### Example Responses
```text
ERROR: call_dispo_report INVALID OR MISSING CAMPAIGNS, INGROUPS, OR DIDS
ERROR: call_dispo_report USER DOES NOT HAVE PERMISSION TO USE THIS FUNCTION
ERROR: call_dispo_report USER DOES NOT HAVE PERMISSION TO VIEW STATS
```

##### Successful Output Format (Summary)
| CAMPAIGN | TOTAL CALLS |
| :--- | :--- |
| TEST_IN2 | 10 |
| TEST_IN3 | 110 |
| TOTAL | 120 |

##### Successful Output Format (with `status_breakdown=1`)
| CAMPAIGN | TOTAL CALLS | ACFLTR | NP | WAITTO | CLOSOP | DROP | INCALL | RQXFER | TD | TESTAM |
| :--- | :--- | :--- | :--- | :--- | :--- | :--- | :--- | :--- | :--- | :--- |
| TEST_IN2 | 10 | 1 | 0 | 0 | 0 | 5 | 0 | 0 | 0 | 4 |
| TEST_IN3 | 110 | 0 | 5 | 13 | 9 | 3 | 1 | 47 | 1 | 31 |
| TOTAL | 120 | 1 | 5 | 13 | 9 | 8 | 1 | 47 | 1 | 35 |

##### Successful Output Format (with `status_breakdown=1` and `show_percentages=1`)
| CAMPAIGN | TOTAL CALLS | ACFLTR | NP | WAITTO | CLOSOP | DROP | INCALL | RQXFER | TD | TESTAM |
| :--- | :--- | :--- | :--- | :--- | :--- | :--- | :--- | :--- | :--- | :--- |
| TEST_IN2 | 10 | 1 (10.0%) | 0 (0.0%) | 0 (0.0%) | 0 (0.0%) | 5 (50.0%) | 0 (0.0%) | 0 (0.0%) | 0 (0.0%) | 4 (40.0%) |
| TEST_IN3 | 110 | 0 (0.0%) | 5 (4.5%) | 13 (11.8%) | 9 (8.2%) | 3 (2.7%) | 1 (0.9%) | 47 (42.7%) | 1 (0.9%) | 31 (28.2%) |
| TOTAL | 120 | 1 (0.8%) | 5 (4.2%) | 13 (10.8%) | 9 (7.5%) | 8 (6.7%) | 1 (0.8%) | 47 (39.2%) | 1 (0.8%) | 35 (29.2%) |

---

### update_campaign
Updates campaign settings in the `vicidial_campaigns` table.

> [!NOTE]
> API user for this function must have `user_level` set to 8 or higher and "modify campaigns" enabled.
> * Avoid special characters like apostrophes, quotes, or ampersands.
> * For optional fields, set them to `--BLANK--` to clear their values.

#### Required Fields
*   `campaign_id` - 2-8 characters
*   `source` - description of what originated the API call (maximum 20 characters)

#### Editable Fields
*   `campaign_name` - 6-40 characters
*   `active` - `Y` or `N`
*   `auto_dial_level` - 2-5 characters, must be within valid range
*   `adaptive_maximum_level` - 3-20 digits
*   `campaign_vdad_exten` - 3-20 characters
*   `hopper_level` - from 1 to 2000
*   `reset_hopper` - `Y` or `N` (default `N`)
*   `dial_method` - `MANUAL`, `RATIO`, `INBOUND_MAN`, `ADAPT_AVERAGE`, `ADAPT_HARD_LIMIT`, or `ADAPT_TAPERED`
*   `dial_timeout` - from 1 to 120
*   `list_order` - order dialed (MUST USE PROPER CASE, e.g.: `UP`, `DOWN`, `RANDOM 2nd NEW`, `DOWN COUNT`)
*   `list_order_randomize` - randomize hopper load: `Y` or `N`
*   `list_order_secondary` - secondary order (MUST USE PROPER CASE, e.g.: `LEAD_ASCEND`, `CALLTIME_DESCEND`)
*   `outbound_cid` - 1-20 digits
*   `dial_status_add` - 1-6 characters
*   `dial_status_remove` - 1-6 characters
*   `lead_filter_id` - valid filter ID, or `---NONE---`
*   `xferconf_one` - `xferconf_five` - Transfer-Conf Overrides: 1-50 characters
*   `dispo_call_url` - Dispo URL, must be URL-encoded (e.g. `http%3A%2F%2Fvicidial.org%2Fquery.php%3Fvar1%3DX%26var2%3Dy`) or set to `'ALT'`
*   `webform_one` - `webform_three` - URL for agent screen webform buttons, must be URL-encoded

#### Example URLs
```http
http://server/vicidial/non_agent_api.php?source=test&user=6666&pass=1234&function=update_campaign&campaign_id=TESTOUT&campaign_name=Updated+test+API+campaign&auto_dial_level=3.0
http://server/vicidial/non_agent_api.php?source=test&user=6666&pass=1234&function=update_campaign&campaign_id=TESTOUT&active=N
http://server/vicidial/non_agent_api.php?source=test&user=6666&pass=1234&function=update_campaign&campaign_id=TESTOUT&reset_hopper=Y
http://server/vicidial/non_agent_api.php?source=test&user=6666&pass=1234&function=update_campaign&campaign_id=TESTOUT&dispo_call_url=http%3A%2F%2Fvicidial.org%2Fquery.php%3Fvar1%3DX%26var2%3Dy
http://server/vicidial/non_agent_api.php?source=test&user=6666&pass=1234&function=update_campaign&campaign_id=TESTOUT&dispo_call_url=ALT
http://server/vicidial/non_agent_api.php?source=test&user=6666&pass=1234&function=update_campaign&campaign_id=TESTOUT&dial_status_add=A2&dial_status_remove=A3&list_order=DOWN+COUNT
```

#### Example Responses
```text
ERROR: update_campaign USER DOES NOT HAVE PERMISSION TO UPDATE CAMPAIGNS - 6666
ERROR: update_campaign YOU MUST USE ALL REQUIRED FIELDS - 6666|1
ERROR: update_campaign NOT AN ALLOWED CAMPAIGN ID - 98762
ERROR: update_campaign CAMPAIGN DOES NOT EXIST - 6666|1000
ERROR: update_campaign AUTO DIAL LEVEL MUST BE FROM 0 TO 18, THIS IS AN OPTIONAL FIELD - 6666|012
ERROR: update_campaign DIAL METHOD IS NOT VALID, THIS IS AN OPTIONAL FIELD - 6666|12
ERROR: update_campaign ADAPT MAXIMUM DIAL LEVEL MUST BE FROM 1 TO 5 CHARACTERS, THIS IS AN OPTIONAL FIELD - 6666|12
ERROR: update_campaign ROUTING EXTENSION MUST BE FROM 3 TO 20 CHARACTERS, THIS IS AN OPTIONAL FIELD - 6666|12
ERROR: update_campaign HOPPER LEVEL MUST BE FROM 1 TO 2000, THIS IS AN OPTIONAL FIELD - 6666|12345
ERROR: update_campaign DIAL TIMEOUT MUST BE FROM 1 TO 120, THIS IS AN OPTIONAL FIELD - 6666|123
ERROR: update_campaign CAMPAIGN NAME MUST BE FROM 6 TO 40 CHARACTERS, THIS IS AN OPTIONAL FIELD - 6666|test
ERROR: update_campaign CAMPAIGN CID MUST BE FROM 6 TO 20 CHARACTERS, THIS IS AN OPTIONAL FIELD - 6666|test
ERROR: update_campaign CAMPAIGN LEAD FILTER MUST BE A VALID FILTER IN THE SYSTEM, THIS IS AN OPTIONAL FIELD - 6666|test
ERROR: update_campaign TRANSFER CONF NUMBER ONE IS NOT VALID, THIS IS AN OPTIONAL FIELD - 6666|123456789012345678901234567890123456789012345678901
ERROR: update_campaign TRANSFER CONF NUMBER TWO IS NOT VALID, THIS IS AN OPTIONAL FIELD - 6666|123456789012345678901234567890123456789012345678901
ERROR: update_campaign TRANSFER CONF NUMBER THREE IS NOT VALID, THIS IS AN OPTIONAL FIELD - 6666|123456789012345678901234567890123456789012345678901
ERROR: update_campaign TRANSFER CONF NUMBER FOUR IS NOT VALID, THIS IS AN OPTIONAL FIELD - 6666|123456789012345678901234567890123456789012345678901
ERROR: update_campaign TRANSFER CONF NUMBER FIVE IS NOT VALID, THIS IS AN OPTIONAL FIELD - 6666|123456789012345678901234567890123456789012345678901
ERROR: update_campaign ACTIVE MUST BE Y OR N, THIS IS AN OPTIONAL FIELD - 6666|U
ERROR: update_campaign LIST ORDER INCLUDING RANDOM ARE NOT ALLOWED, THIS IS AN OPTIONAL FIELD - |6666|RANDOM 2nd NEW|
ERROR: update_campaign LIST ORDER IS NOT VALID, THIS IS AN OPTIONAL FIELD - |6666|Y|
ERROR: update_campaign LIST ORDER IS NOT A VALID LENGTH, THIS IS AN OPTIONAL FIELD - |6666|Y|
ERROR: update_campaign LIST ORDER RANDOMIZE IS NOT VALID, THIS IS AN OPTIONAL FIELD - |6666|X|
ERROR: update_campaign LIST ORDER RANDOMIZE IS NOT A VALID LENGTH, THIS IS AN OPTIONAL FIELD - |6666|YES|
ERROR: update_campaign LIST ORDER SECONDARY IS NOT VALID, THIS IS AN OPTIONAL FIELD - |6666|Y|
ERROR: update_campaign LIST ORDER SECONDARY IS NOT A VALID LENGTH, THIS IS AN OPTIONAL FIELD - |6666|LEAD|
ERROR: update_campaign DIAL STATUS ADD IS NOT VALID, THIS IS AN OPTIONAL FIELD - |6666|A1234567|
ERROR: update_campaign DIAL STATUS ADD FAILURE, TOO MANY EXISTING DIAL STATUSES, THIS IS AN OPTIONAL FIELD - |6666|A12345|
ERROR: update_campaign DIAL STATUS REMOVE IS NOT VALID, THIS IS AN OPTIONAL FIELD - |6666|A1234567|
ERROR: update_campaign DIAL STATUS REMOVE FAILURE, NO EXISTING DIAL STATUSES SET, THIS IS AN OPTIONAL FIELD - |666666|A12345|
ERROR: update_campaign DIAL STATUS REMOVE FAILURE, NOT AN EXISTING DIAL STATUS, THIS IS AN OPTIONAL FIELD - |666666|A12345|
ERROR: update_campaign DISPO CALL URL IS NOT VALID, THIS IS AN OPTIONAL FIELD: |6666|XY
ERROR: update_campaign WENFORM 1 URL IS NOT VALID, THIS IS AN OPTIONAL FIELD: |6666|XY
ERROR: update_campaign WENFORM 2 URL IS NOT VALID, THIS IS AN OPTIONAL FIELD: |6666|XY
ERROR: update_campaign WENFORM 3 URL IS NOT VALID, THIS IS AN OPTIONAL FIELD: |6666|XY
NOTICE: update_campaign NO UPDATES DEFINED - 6666|
NOTICE: update_campaign HOPPER HAS BEEN RESET - 6666|1101|324
SUCCESS: update_campaign CAMPAIGN HAS BEEN UPDATED - 6666|1101
```

---

### update_alt_url
Updates, adds, or displays alternate dispo call URL entries for a campaign.

> [!NOTE]
> API user for this function must have `user_level` set to 8 or higher and "modify campaigns" enabled.
> * Avoid special characters like apostrophes, quotes, or ampersands.
> * Send `--BLANK--` as the value to clear `url_description` or `url_lists`.

#### Required Fields
*   `campaign_id` - 2-8 characters
*   `source` - description of what originated the API call (maximum 20 characters)
*   `entry_type` - `'campaign'`
*   `url_type` - `dispo`, `start`, `addlead`, or `noagent`
*   `alt_url_id` - digits only, `'NEW'` or `'LIST'`. Specifies the ID of the URL to edit or `NEW` to create one.
    *   *If only one Alt URL exists, this can be left blank.*
    *   *Settings fields `stage` (`csv`, `tab`, `pipe`) and `header` (`YES`, `NO`) are supported when `alt_url_id` is set to `LIST`.*

#### Editable Fields
*   `active` - `Y` or `N` (default `N`)
*   `url_rank` - digits only
*   `url_statuses` - *For DISPO URLs only.* Space-separated list of active statuses (e.g. `"SALE1 SALE2"`) or `"---ALL---"`
*   `url_description` - text description of the URL
*   `url_lists` - space-separated list IDs (e.g. `"101 102"`) or blank `""` for all lists
*   `url_call_length` - digits only, minimum call duration (seconds), default is 0
*   `url_address` - URL for the Alt URL, must be URL-encoded

#### Example URLs
```http
http://server/vicidial/non_agent_api.php?source=test&user=6666&pass=1234&function=update_alt_url&campaign_id=ALTTEST&entry_type=campaign&url_type=dispo&alt_url_id=3&active=N
http://server/vicidial/non_agent_api.php?source=test&user=6666&pass=1234&function=update_alt_url&campaign_id=ALTTEST&entry_type=campaign&url_type=dispo&alt_url_id=3&url_rank=2
http://server/vicidial/non_agent_api.php?source=test&user=6666&pass=1234&function=update_alt_url&campaign_id=ALTTEST&entry_type=campaign&url_type=dispo&alt_url_id=3&url_statuses=SALE+PSALE
http://server/vicidial/non_agent_api.php?source=test&user=6666&pass=1234&function=update_alt_url&campaign_id=ALTTEST&entry_type=campaign&url_type=dispo&alt_url_id=3&url_description=Updated+SALE+URL
http://server/vicidial/non_agent_api.php?source=test&user=6666&pass=1234&function=update_alt_url&campaign_id=ALTTEST&entry_type=campaign&url_type=dispo&alt_url_id=3&url_lists=101+102
http://server/vicidial/non_agent_api.php?source=test&user=6666&pass=1234&function=update_alt_url&campaign_id=ALTTEST&entry_type=campaign&url_type=dispo&alt_url_id=3&url_call_length=2
http://server/vicidial/non_agent_api.php?source=test&user=6666&pass=1234&function=update_alt_url&campaign_id=ALTTEST&entry_type=campaign&url_type=dispo&alt_url_id=3&url_address=http%3A%2F%2Fvicidial.org%2Fquery.php%3Fvar1%3DX%26var2%3Dy
http://server/vicidial/non_agent_api.php?source=test&user=6666&pass=1234&function=update_alt_url&campaign_id=ALTTEST&entry_type=campaign&url_type=dispo&alt_url_id=LIST
http://server/vicidial/non_agent_api.php?source=test&user=6666&pass=1234&function=update_alt_url&campaign_id=ALTTEST&entry_type=campaign&url_type=dispo&alt_url_id=NEW&url_statuses=PRQUAL&url_description=Testing&url_address=http%3A%2F%2Fvicidial.org%2Fquery.php%3Fvar1%3DX%26var2%3Dy
```

#### Example Responses
```text
ERROR: update_alt_url USER DOES NOT HAVE PERMISSION TO UPDATE CAMPAIGNS - 6666
ERROR: update_alt_url YOU MUST USE ALL REQUIRED FIELDS - 6666|1
ERROR: update_alt_url NOT AN ALLOWED CAMPAIGN ID - 98762
ERROR: update_alt_url CAMPAIGN DOES NOT EXIST - 6666|1000
ERROR: update_alt_url NO VALID URL TYPE DEFINED: - 6666|||
ERROR: update_alt_url NO ENTRY TYPE DEFINED - 6666||
ERROR: update_alt_url CAMPAIGNS dispo URL IS NOT SET TO ALT - 6666|dispo|campaign|TESTCAMP
ERROR: update_alt_url ALT URL ID DOES NOT EXIST - 6666|2|dispo|campaign|TESTCAMP|2
ERROR: update_alt_url ACTIVE MUST BE Y OR N, THIS IS AN OPTIONAL FIELD - 6666|U
ERROR: update_alt_url URL ADDRESS IS NOT VALID, THIS IS AN OPTIONAL FIELD - 6666|
ERROR: update_alt_url URL RANK IS NOT VALID, THIS IS AN OPTIONAL FIELD - 6666|999999
ERROR: update_alt_url URL CAL LENGTH IS NOT VALID, THIS IS AN OPTIONAL FIELD - 6666|999999
ERROR: update_alt_url URL STATUSES IS NOT VALID, THIS IS AN OPTIONAL FIELD - 6666|
ERROR: update_alt_url URL DESCRIPTION IS NOT VALID, THIS IS AN OPTIONAL FIELD - 6666|
ERROR: update_alt_url URL LISTS IS NOT VALID, THIS IS AN OPTIONAL FIELD - 6666|
NOTICE: update_campaign NO UPDATES DEFINED - 6666|
SUCCESS: update_alt_url ALT URL HAS BEEN UPDATED - 6666|1|dispo|campaign|TESTCAMP
SUCCESS: update_alt_url ALT URL HAS BEEN ADDED - 6666|NEW URL ID: 2|dispo|campaign|TESTCAMP
NOTICE: update_alt_url LIST, No Records Found - 6666|dispo|campaign|TESTCAMP|0
```

##### LIST Response Output Format
| url_id | campaign_id | entry_type | active | url_type | url_rank | url_statuses | url_description | url_lists | url_call_length | url_address |
| :--- | :--- | :--- | :--- | :--- | :--- | :--- | :--- | :--- | :--- | :--- |
| 3 | ALTTEST | campaign | Y | dispo | 2 | SALE PSALE | Updated SALE URL | | 2 | `http://vicidial.org/query.php?var1=X&var2=y` |

---

### update_presets
Updates, adds, displays, or deletes campaign preset entries.

> [!NOTE]
> API user for this function must have `user_level` set to 8 or higher and "modify campaigns" enabled.
> * Avoid special characters like apostrophes, quotes, or ampersands.
> * Send `--BLANK--` as the value to clear `preset_dtmf`.

#### Required Fields
*   `campaign_id` - 2-8 characters
*   `source` - description of what originated the API call (maximum 20 characters)
*   `preset_name` - name of the preset. *Optional if only one preset exists.*
*   `action` - `UPDATE`, `NEW`, `DELETE` or `LIST` (default is `UPDATE`)
    *   *Settings fields `stage` (`csv`, `tab`, `pipe` [default]) and `header` (`YES`, `NO`) are supported when action is `LIST`.*

#### Editable Fields
*   `preset_hide_number` - `Y` or `N` (default `N`)
*   `preset_number` - digits only (1-50 digits)
*   `preset_dtmf` - digits and certain characters only (`0123456789PSQ`, where `P` = #, `S` = \*, `Q` = half-second pause)

#### Example URLs
```http
http://server/vicidial/non_agent_api.php?source=test&user=6666&pass=1234&function=update_presets&campaign_id=ALTTEST&preset_name=SALES&preset_number=3125551212&preset_hide_number=N&action=NEW
http://server/vicidial/non_agent_api.php?source=test&user=6666&pass=1234&function=update_presets&campaign_id=ALTTEST&preset_number=3125551213&preset_hide_number=N&preset_dtmf=1234
http://server/vicidial/non_agent_api.php?source=test&user=6666&pass=1234&function=update_presets&campaign_id=ALTTEST&action=LIST
http://server/vicidial/non_agent_api.php?source=test&user=6666&pass=1234&function=update_presets&campaign_id=ALTTEST&preset_name=SUPPORT+LINE&preset_number=3125551214&preset_hide_number=Y&preset_dtmf=1&action=NEW
http://server/vicidial/non_agent_api.php?source=test&user=6666&pass=1234&function=update_presets&campaign_id=ALTTEST&preset_name=SUPPORT+LINE&action=DELETE
```

#### Example Responses
```text
ERROR: update_presets USER DOES NOT HAVE PERMISSION TO UPDATE CAMPAIGNS - 6666
ERROR: update_presets YOU MUST USE ALL REQUIRED FIELDS - 6666|1
ERROR: update_presets NOT AN ALLOWED CAMPAIGN ID - 98762
ERROR: update_presets CAMPAIGN DOES NOT EXIST - 6666|1000
ERROR: update_presets PRESET NAME DOES NOT EXIST - 6666|2|dispo|campaign|TESTCAMP|2
ERROR: update_presets PRESET HIDE NUMBER MUST BE Y OR N, THIS IS AN OPTIONAL FIELD - 6666|U
ERROR: update_presets PRESET DTMF IS NOT VALID, THIS IS AN OPTIONAL FIELD - 6666|
NOTICE: update_campaign NO UPDATES DEFINED - 6666|
SUCCESS: update_presets PRESET HAS BEEN UPDATED - 6666|1|campaign|TESTCAMP
SUCCESS: update_presets PRESET HAS BEEN ADDED - 6666|NEW PRESET: SUPPORT LINE|TESTCAMP
SUCCESS: update_presets PRESET HAS BEEN DELETED - 6666|SUPPORT LINE|TESTCAMP
NOTICE: update_presets LIST, No Records Found - 6666|TESTCAMP|0
```

##### LIST Response Output Format
| preset_name | campaign_id | preset_number | preset_hide_number | preset_dtmf |
| :--- | :--- | :--- | :--- | :--- |
| SALES | ALTTEST | 3125551213 | N | 1234 |

---

### add_did
Adds new Inbound DID entries to the `vicidial_inbound_dids` table.

> [!NOTE]
> API user for this function must have `user_level` set to 8 or higher and "Modify DIDs" enabled.
> Special characters in the `did_pattern` must be URL encoded (e.g. `+` must be sent as `%2B`).

#### Required Fields
*   `did_pattern` - 2-50 characters
*   `source` - description of what originated the API call (maximum 20 characters)

#### Editable Fields
*   `did_description` - 6-50 characters
*   `active` - `Y` or `N`
*   `did_route` - `EXTEN`, `VOICEMAIL`, `AGENT`, `PHONE`, `IN_GROUP`, `CALLMENU`, or `VMAIL_NO_INST`
*   `record_call` - `Y`, `N`, or `Y_QUEUESTOP`
*   `extension` - 1-50 digits
*   `exten_context` - 1-50 characters
*   `voicemail_ext` - 1-10 digits
*   `phone_extension` - 1-100 characters
*   `server_ip` - must be valid IP address (maximum 15 characters)
*   `group` - valid inbound group ID
*   `menu_id` - valid call menu ID
*   `filter_clean_cid_number` - 1-20 characters
*   `call_handle_method` - for `IN_GROUP` route (`CID`, `CIDLOOKUP`, `CLOSER`, etc., default `'CID'`)
*   `agent_search_method` - for `IN_GROUP` route (`LO` = Load-Balanced-Overflow, `LB` = Load-Balanced, `SO` = Server-Only, default `'LB'`)
*   `list_id` - for `IN_GROUP` route (3-12 digits, default `'999'`)
*   `entry_list_id` - for `IN_GROUP` route (1-12 digits, default `'0'`)
*   `campaign_id` - for `IN_GROUP` route (default `''`)
*   `phone_code` - for `IN_GROUP` route (default `'1'`)

#### Example URL
```http
http://server/vicidial/non_agent_api.php?source=test&user=6666&pass=1234&function=add_did&did_pattern=%2B1727555112&did_description=Updated+test+API+DID&active=Y
http://server/vicidial/non_agent_api.php?source=test&user=6666&pass=1234&function=add_did&did_pattern=7275553331&did_description=Testing+DID&active=N&did_route=EXTEN&record_call=Y_QUEUESTOP&extension=8300&exten_context=notdefault&voicemail_ext=8301&phone_extension=55000&server_ip=192.168.198.5&group=TEST_IN2&filter_clean_cid_number=R10
```

#### Example Responses
```text
ERROR: add_did USER DOES NOT HAVE PERMISSION TO ADD DIDS - 6666
ERROR: add_did YOU MUST USE ALL REQUIRED FIELDS - 6666|1
ERROR: add_did NOT AN ALLOWED DID - 98762
ERROR: add_did DID ALREADY EXISTS - 6666|1000
ERROR: add_did DID DESCRIPTION MUST BE FROM 6 TO 50 CHARACTERS, THIS IS AN OPTIONAL FIELD - 6666|012
ERROR: add_did ACTIVE MUST BE Y OR N, THIS IS AN OPTIONAL FIELD - 6666|U
ERROR: add_did DID ROUTE IS NOT VALID, THIS IS AN OPTIONAL FIELD - 6666|12
ERROR: add_did RECORD CALL MUST BE Y OR N, THIS IS AN OPTIONAL FIELD - 6666|
ERROR: add_did EXTENSION MUST BE FROM 1 TO 50 CHARACTERS, THIS IS AN OPTIONAL FIELD - 6666|
ERROR: add_did EXTENSION CONTEXT MUST BE FROM 1 TO 50 CHARACTERS, THIS IS AN OPTIONAL FIELD - 6666|
ERROR: add_did VOICEMAIL EXTENSION MUST BE FROM 1 TO 10 CHARACTERS, THIS IS AN OPTIONAL FIELD - 6666|
ERROR: add_did PHONE EXTENSION MUST BE FROM 1 TO 100 CHARACTERS, THIS IS AN OPTIONAL FIELD - 6666|
ERROR: add_did SERVER IP MUST BE A VALID SERVER IN THE SYSTEM, THIS IS AN OPTIONAL FIELD - 6666|0.0.0.0
ERROR: add_did GROUP ID MUST BE A VALID INBOUND GROUP IN THE SYSTEM, THIS IS AN OPTIONAL FIELD - 6666|test
ERROR: add_did MENU ID MUST BE A VALID CALL MENU IN THE SYSTEM, THIS IS AN OPTIONAL FIELD - 6666|test_menu
ERROR: add_did CLEAN CID NUMBER MUST BE FROM 1 TO 20 CHARACTERS, THIS IS AN OPTIONAL FIELD - 6666|
ERROR: add_did CALL HANDLE METHOD MUST BE FROM 3 TO 20 CHARACTERS, THIS IS AN OPTIONAL FIELD - 6666|1
ERROR: add_did AGENT SEARCH METHOD MUST BE 2 CHARACTERS, THIS IS AN OPTIONAL FIELD - 6666|1
ERROR: add_did LIST ID MUST BE FROM 3 TO 14 DIGITS, THIS IS AN OPTIONAL FIELD - 6666|1
ERROR: add_did ENTRY LIST ID MUST BE FROM 1 TO 14 DIGITS, THIS IS AN OPTIONAL FIELD - 6666|999999999999999
ERROR: add_did CAMPAIGN MUST BE FROM 1 TO 8 CHARACTERS, THIS IS AN OPTIONAL FIELD - 6666|TESTCAMPS
ERROR: add_did PHONE CODE MUST BE FROM 1 TO 10 CHARACTERS, THIS IS AN OPTIONAL FIELD - 6666|99999999999
SUCCESS: add_did DID HAS BEEN ADDED - 6666|1101
```

---

### copy_did
Adds new Inbound DID entries, copied from an existing DID entry.

> [!NOTE]
> API user for this function must have `user_level` set to 8 or higher and "Modify DIDs" enabled.
> Special characters in patterns must be URL encoded (e.g. `+` must be sent as `%2B`).

#### Required Fields
*   `source_did_pattern` - 2-50 characters, must be a valid DID in the system
*   `new_dids` - 2-2000 characters, multiple DIDs separated by commas
*   `source` - description of what originated the API call (maximum 20 characters)

#### Optional Fields
*   `did_description` - 6-50 characters. If left blank, the source DID description will be used.

#### Example URL
```http
http://server/vicidial/non_agent_api.php?source=test&user=6666&pass=1234&function=copy_did&source_did_pattern=%2B1727555112&did_description=Updated+test+API+DID&new_dids=%2B1727555119
http://server/vicidial/non_agent_api.php?source=test&user=6666&pass=1234&function=copy_did&source_did_pattern=7275551114&new_dids=7275553338,7275553337,7275553336
```

#### Example Responses
```text
ERROR: copy_did USER DOES NOT HAVE PERMISSION TO ADD DIDS - 6666
ERROR: copy_did YOU MUST USE ALL REQUIRED FIELDS - 6666|1|1
ERROR: copy_did NOT AN ALLOWED DID - 98762
ERROR: copy_did DID DESCRIPTION MUST BE EITHER BLANK OR FROM 6 TO 50 CHARACTERS, THIS IS AN OPTIONAL FIELD - 6666|012
NOTICE: copy_did DID ALREADY EXISTS - 6666|1000
NOTICE: copy_did DID HAS BEEN ADDED - 6666|1000|1001|12|1
ERROR: copy_did NO DIDS ADDED - 6666|1000|0|1
SUCCESS: copy_did DIDS HAVE BEEN ADDED - 6666|1000|1|0
```

> [!NOTE]
> Detailed output is printed for each attempted DID, followed by a summary success or error message.

---

### update_did
Updates Inbound DID information in the `vicidial_inbound_dids` table.

> [!NOTE]
> API user for this function must have `user_level` set to 8 or higher and "Modify DIDs" enabled.
> Special characters in the `did_pattern` must be URL encoded (e.g. `+` must be sent as `%2B`).

#### Required Fields
*   `did_pattern` - 2-50 characters, must be a valid DID in the system
*   `source` - description of what originated the API call (maximum 20 characters)

#### Editable Fields
*   `did_description` - 6-50 characters
*   `active` - `Y` or `N`
*   `did_route` - `EXTEN`, `VOICEMAIL`, `AGENT`, `PHONE`, `IN_GROUP`, `CALLMENU`, or `VMAIL_NO_INST`
*   `record_call` - `Y`, `N`, or `Y_QUEUESTOP`
*   `extension` - 1-50 digits
*   `exten_context` - 1-50 characters
*   `voicemail_ext` - 1-10 digits
*   `phone_extension` - 1-100 characters
*   `server_ip` - must be valid IP address (maximum 15 characters)
*   `group` - valid inbound group ID
*   `menu_id` - valid call menu ID
*   `filter_clean_cid_number` - 1-20 characters
*   `call_handle_method` - for `IN_GROUP` route (`CID`, `CIDLOOKUP`, `CLOSER`, etc., default `'CID'`)
*   `agent_search_method` - for `IN_GROUP` route (`LO` = Load-Balanced-Overflow, `LB` = Load-Balanced, `SO` = Server-Only, default `'LB'`)
*   `list_id` - for `IN_GROUP` route (3-12 digits, default `'999'`)
*   `entry_list_id` - for `IN_GROUP` route (1-12 digits, default `'0'`)
*   `campaign_id` - for `IN_GROUP` route (default `''`)
*   `phone_code` - for `IN_GROUP` route (default `'1'`)
*   `delete_did` - `Y` or `N` (default `N`)

#### Example URL
```http
http://server/vicidial/non_agent_api.php?source=test&user=6666&pass=1234&function=update_did&did_pattern=%2B1727555112&did_description=Updated+test+API+DID&active=Y
http://server/vicidial/non_agent_api.php?source=test&user=6666&pass=1234&function=update_did&did_pattern=7275553331&did_description=Testing+DID&active=N&did_route=EXTEN&record_call=Y_QUEUESTOP&extension=8300&exten_context=notdefault&voicemail_ext=8301&phone_extension=55000&server_ip=192.168.198.5&group=TEST_IN2&filter_clean_cid_number=R10
```

#### Example Responses
```text
ERROR: update_did USER DOES NOT HAVE PERMISSION TO UPDATE DIDS - 6666
ERROR: update_did USER DOES NOT HAVE PERMISSION TO DELETE DIDS - 6666
ERROR: update_did YOU MUST USE ALL REQUIRED FIELDS - 6666|1
ERROR: update_did NOT AN ALLOWED DID - 98762
ERROR: update_did DID DOES NOT EXIST - 6666|1000
ERROR: update_did DID DESCRIPTION MUST BE FROM 6 TO 50 CHARACTERS, THIS IS AN OPTIONAL FIELD - 6666|012
ERROR: update_did ACTIVE MUST BE Y OR N, THIS IS AN OPTIONAL FIELD - 6666|U
ERROR: update_did DID ROUTE IS NOT VALID, THIS IS AN OPTIONAL FIELD - 6666|12
ERROR: update_did RECORD CALL MUST BE Y OR N, THIS IS AN OPTIONAL FIELD - 6666|
ERROR: update_did EXTENSION MUST BE FROM 1 TO 50 CHARACTERS, THIS IS AN OPTIONAL FIELD - 6666|
ERROR: update_did EXTENSION CONTEXT MUST BE FROM 1 TO 50 CHARACTERS, THIS IS AN OPTIONAL FIELD - 6666|
ERROR: update_did VOICEMAIL EXTENSION MUST BE FROM 1 TO 10 CHARACTERS, THIS IS AN OPTIONAL FIELD - 6666|
ERROR: update_did PHONE EXTENSION MUST BE FROM 1 TO 100 CHARACTERS, THIS IS AN OPTIONAL FIELD - 6666|
ERROR: update_did SERVER IP MUST BE A VALID SERVER IN THE SYSTEM, THIS IS AN OPTIONAL FIELD - 6666|0.0.0.0
ERROR: update_did GROUP ID MUST BE A VALID INBOUND GROUP IN THE SYSTEM, THIS IS AN OPTIONAL FIELD - 6666|test
ERROR: update_did MENU ID MUST BE A VALID CALL MENU IN THE SYSTEM, THIS IS AN OPTIONAL FIELD - 6666|test_menu
ERROR: update_did CLEAN CID NUMBER MUST BE FROM 1 TO 20 CHARACTERS, THIS IS AN OPTIONAL FIELD - 6666|
ERROR: update_did CALL HANDLE METHOD MUST BE FROM 3 TO 20 CHARACTERS, THIS IS AN OPTIONAL FIELD - 6666|1
ERROR: update_did AGENT SEARCH METHOD MUST BE 2 CHARACTERS, THIS IS AN OPTIONAL FIELD - 6666|1
ERROR: update_did LIST ID MUST BE FROM 3 TO 14 DIGITS, THIS IS AN OPTIONAL FIELD - 6666|1
ERROR: update_did ENTRY LIST ID MUST BE FROM 1 TO 14 DIGITS, THIS IS AN OPTIONAL FIELD - 6666|999999999999999
ERROR: update_did CAMPAIGN MUST BE FROM 1 TO 8 CHARACTERS, THIS IS AN OPTIONAL FIELD - 6666|TESTCAMPS
ERROR: update_did PHONE CODE MUST BE FROM 1 TO 10 CHARACTERS, THIS IS AN OPTIONAL FIELD - 6666|99999999999
NOTICE: update_did NO UPDATES DEFINED - 6666|
SUCCESS: update_did DID HAS BEEN DELETED - 6666|23|1101
SUCCESS: update_did DID HAS BEEN UPDATED - 6666|1101
```

---

### update_cid_group_entry
Updates CID Group entries in the `vicidial_campaign_cid_areacodes` table.

> [!NOTE]
> API user for this function must have `user_level` set to 8 or higher and "Modify Campaigns" enabled.

#### Required Fields
*   `cid_group_id` - 2-20 characters, must be a valid CID Group ID or Campaign ID
*   `source` - description of what originated the API call (maximum 20 characters)
*   `areacode` - AREACODE (e.g. `"312"`), STATE (e.g. `"FL"`), `"NONE"`, or `"---ALL---"`
*   `stage` - `UPDATE`, `ADD`, `DELETE` or `INFO`

#### Optional Fields
*   `outbound_cid` - 1-20 digits (or `"---ALL---"` or blank to affect all entries for an areacode)

#### Editable Fields
*   `cid_description` - 6-50 characters (default is blank)
*   `active` - `Y` or `N` (default `N`)

#### Example URLs
```http
http://server/vicidial/non_agent_api.php?source=test&user=6666&pass=1234&function=update_cid_group_entry&stage=INFO&cid_group_id=TESTINX5&areacode=---ALL---
http://server/vicidial/non_agent_api.php?source=test&user=6666&pass=1234&function=update_cid_group_entry&stage=ADD&cid_group_id=testing_here4&areacode=312&outbound_cid=3125551212&cid_description=Testing+DID&active=N
http://server/vicidial/non_agent_api.php?source=test&user=6666&pass=1234&function=update_cid_group_entry&stage=DELETE&cid_group_id=testing_here4&areacode=312&outbound_cid=3125551212
http://server/vicidial/non_agent_api.php?source=test&user=6666&pass=1234&function=update_cid_group_entry&stage=UPDATE&cid_group_id=testing_here4&areacode=312&outbound_cid=3125551212&active=Y
http://server/vicidial/non_agent_api.php?source=test&user=6666&pass=1234&function=update_cid_group_entry&stage=UPDATE&cid_group_id=testing_here4&areacode=312&outbound_cid=---ALL---&active=N&cid_description=Deactivated+for+now
```

#### Example Responses
```text
ERROR: update_cid_group_entry USER DOES NOT HAVE PERMISSION TO UPDATE CID GROUPS - 6666
ERROR: update_cid_group_entry YOU MUST USE ALL REQUIRED FIELDS - 6666|1||
ERROR: update_cid_group_entry NOT AN ALLOWED CID GROUP - 98762
ERROR: update_cid_group_entry CID GROUP DOES NOT EXIST - 6666|1000
ERROR: update_cid_group_entry THIS CID GROUP HAS NO ENTRIES - 6666|1000
```

##### Successful INFO Response Output Format
```text
SUCCESS: update_cid_group_entry CID GROUP INFO AS FOLLOWS - 1101|---ALL---|---ALL---|3
1|727|7275551212|N|FL1 test|0
2|727|7275551313|N|FL1 test|0
4|998|7275551114|N|FL2 test|0
```

##### Successful Modify Responses
```text
ERROR: update_cid_group_entry YOU MUST DEFINE AN AREACODE WHEN ADDING AND ENTRY - 6666|1000|---ALL---
ERROR: update_cid_group_entry YOU MUST DEFINE A CID NUMBER WHEN ADDING AND ENTRY - 6666|1000|1
ERROR: update_cid_group_entry THIS CID GROUP ENTRY ALREADY EXISTS - 6666|1000|813|8135551212
SUCCESS: update_cid_group_entry CID GROUP ENTRY HAS BEEN ADDED - 6666|1101|813|8135551212|1
SUCCESS: update_cid_group_entry CID GROUP ENTRIES HAVE BEEN DELETED - 6666|1101|813|8135551212|1
ERROR: update_cid_group_entry NO UPDATES DEFINED - 6666|813|8135551212||
SUCCESS: update_cid_group_entry CID GROUP ENTRIES HAVE BEEN UPDATED - 6666|1101|813|8135551212|1
```