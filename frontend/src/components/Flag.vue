<script setup lang="ts">
defineProps({
  country: String,
})
</script>

/* cSpell:disable */
<template>
  <icon-flagpack-ad v-if="country === 'Andorra'" />
  <icon-flagpack-ae v-else-if="country === 'United Arab Emirates'" />
  <icon-flagpack-af v-else-if="country === 'Afghanistan'" />
  <icon-flagpack-ag v-else-if="country === 'Antigua and Barbuda'" />
  <icon-flagpack-ai v-else-if="country === 'Anguilla'" />
  <icon-flagpack-al v-else-if="country === 'Albania'" />
  <icon-flagpack-am v-else-if="country === 'Armenia'" />
  <icon-flagpack-ao v-else-if="country === 'Angola'" />
  <icon-flagpack-aq v-else-if="country === 'Antarctica'" />
  <icon-flagpack-ar v-else-if="country === 'Argentina'" />
  <icon-flagpack-as v-else-if="country === 'American Samoa'" />
  <icon-flagpack-at v-else-if="country === 'Austria'" />
  <icon-flagpack-au v-else-if="country === 'Australia'" />
  <icon-flagpack-aw v-else-if="country === 'Aruba'" />
  <icon-flagpack-ax v-else-if="country === 'Åland Islands'" />
  <icon-flagpack-az v-else-if="country === 'Azerbaijan'" />
  <icon-flagpack-ba v-else-if="country === 'Bosnia and Herzegovina'" />
  <icon-flagpack-bb v-else-if="country === 'Barbados'" />
  <icon-flagpack-bd v-else-if="country === 'Bangladesh'" />
  <icon-flagpack-be v-else-if="country === 'Belgium'" />
  <icon-flagpack-bf v-else-if="country === 'Burkina Faso'" />
  <icon-flagpack-bg v-else-if="country === 'Bulgaria'" />
  <icon-flagpack-bh v-else-if="country === 'Bahrain'" />
  <icon-flagpack-bi v-else-if="country === 'Burundi'" />
  <icon-flagpack-bj v-else-if="country === 'Benin'" />
  <icon-flagpack-bl v-else-if="country === 'Saint Barthélemy'" />
  <icon-flagpack-bm v-else-if="country === 'Bermuda'" />
  <icon-flagpack-bn v-else-if="country === 'Brunei Darussalam'" />
  <icon-flagpack-bo v-else-if="country === 'Bolivia'" />
  <icon-flagpack-bq-bo v-else-if="country === 'Bonaire'" />
  <icon-flagpack-bq-sa v-else-if="country === 'Saba'" />
  <icon-flagpack-bq-se v-else-if="country === 'Sint Eustatius'" />
  <icon-flagpack-br v-else-if="country === 'Brazil'" />
  <icon-flagpack-bs v-else-if="country === 'Bahamas'" />
  <icon-flagpack-bt v-else-if="country === 'Bhutan'" />
  <icon-flagpack-bv v-else-if="country === 'Bouvet Island'" />
  <icon-flagpack-bw v-else-if="country === 'Botswana'" />
  <icon-flagpack-by v-else-if="country === 'Belarus'" />
  <icon-flagpack-bz v-else-if="country === 'Belize'" />
  <icon-flagpack-ca v-else-if="country === 'Canada'" />
  <icon-flagpack-cc v-else-if="country === 'Cocos Islands'" />
  <icon-flagpack-cd v-else-if="country === 'Democratic Republic of the Congo'" />
  <icon-flagpack-cf v-else-if="country === 'Central African Republic'" />
  <icon-flagpack-cg v-else-if="country === 'Congo'" />
  <icon-flagpack-ch v-else-if="country === 'Switzerland'" />
  <icon-flagpack-ci v-else-if="country === 'Côte d\'Ivoire'" />
  <icon-flagpack-ck v-else-if="country === 'Cook Islands'" />
  <icon-flagpack-cl v-else-if="country === 'Chile'" />
  <icon-flagpack-cm v-else-if="country === 'Cameroon'" />
  <icon-flagpack-cn v-else-if="country === 'China'" />
  <icon-flagpack-co v-else-if="country === 'Colombia'" />
  <icon-flagpack-cr v-else-if="country === 'Costa Rica'" />
  <icon-flagpack-cu v-else-if="country === 'Cuba'" />
  <icon-flagpack-cv v-else-if="country === 'Cabo Verde'" />
  <icon-flagpack-cw v-else-if="country === 'Curaçao'" />
  <icon-flagpack-cx v-else-if="country === 'Christmas Island'" />
  <icon-flagpack-cy v-else-if="country === 'Cyprus'" />
  <icon-flagpack-cz v-else-if="country === 'Czechia'" />
  <icon-flagpack-de v-else-if="country === 'Germany'" />
  <icon-flagpack-dj v-else-if="country === 'Djibouti'" />
  <icon-flagpack-dk v-else-if="country === 'Denmark'" />
  <icon-flagpack-dm v-else-if="country === 'Dominica'" />
  <icon-flagpack-do v-else-if="country === 'Dominican Republic'" />
  <icon-flagpack-dz v-else-if="country === 'Algeria'" />
  <icon-flagpack-ec v-else-if="country === 'Ecuador'" />
  <icon-flagpack-ee v-else-if="country === 'Estonia'" />
  <icon-flagpack-eg v-else-if="country === 'Egypt'" />
  <icon-flagpack-eh v-else-if="country === 'Western Sahara'" />
  <icon-flagpack-er v-else-if="country === 'Eritrea'" />
  <icon-flagpack-es v-else-if="country === 'Spain'" />
  <icon-flagpack-et v-else-if="country === 'Ethiopia'" />
  <icon-flagpack-fi v-else-if="country === 'Finland'" />
  <icon-flagpack-fj v-else-if="country === 'Fiji'" />
  <icon-flagpack-fk v-else-if="country === 'Falkland Islands'" />
  <icon-flagpack-fm v-else-if="country === 'Micronesia'" />
  <icon-flagpack-fo v-else-if="country === 'Faroe Islands'" />
  <icon-flagpack-fr v-else-if="country === 'France'" />
  <icon-flagpack-ga v-else-if="country === 'Gabon'" />
  <icon-flagpack-gb v-else-if="country === 'England'" />
  <icon-flagpack-gb-sct v-else-if="country === 'Scotland'" />
  <icon-flagpack-gb-ukm v-else-if="country === 'United Kingdom of Great Britain'" />
  <icon-flagpack-gb-wls v-else-if="country === 'Wales'" />
  <icon-flagpack-gd v-else-if="country === 'Grenada'" />
  <icon-flagpack-ge v-else-if="country === 'Georgia'" />
  <icon-flagpack-gf v-else-if="country === 'French Guiana'" />
  <icon-flagpack-gg v-else-if="country === 'Guernsey'" />
  <icon-flagpack-gh v-else-if="country === 'Ghana'" />
  <icon-flagpack-gi v-else-if="country === 'Gibraltar'" />
  <icon-flagpack-gl v-else-if="country === 'Greenland'" />
  <icon-flagpack-gm v-else-if="country === 'Gambia'" />
  <icon-flagpack-gn v-else-if="country === 'Guinea'" />
  <icon-flagpack-gp v-else-if="country === 'Guadeloupe'" />
  <icon-flagpack-gq v-else-if="country === 'Equatorial Guinea'" />
  <icon-flagpack-gr v-else-if="country === 'Greece'" />
  <icon-flagpack-gs v-else-if="country === 'South Georgia and the South Sandwich Islands'" />
  <icon-flagpack-gt v-else-if="country === 'Guatemala'" />
  <icon-flagpack-gu v-else-if="country === 'Guam'" />
  <icon-flagpack-gw v-else-if="country === 'Guinea-Bissau'" />
  <icon-flagpack-gy v-else-if="country === 'Guyana'" />
  <icon-flagpack-hk v-else-if="country === 'Hong Kong'" />
  <icon-flagpack-hm v-else-if="country === 'Heard Island and McDonald Islands'" />
  <icon-flagpack-hn v-else-if="country === 'Honduras'" />
  <icon-flagpack-hr v-else-if="country === 'Croatia'" />
  <icon-flagpack-ht v-else-if="country === 'Haiti'" />
  <icon-flagpack-hu v-else-if="country === 'Hungary'" />
  <icon-flagpack-id v-else-if="country === 'Indonesia'" />
  <icon-flagpack-ie v-else-if="country === 'Ireland'" />
  <icon-flagpack-il v-else-if="country === 'Isreal'" />
  <icon-flagpack-im v-else-if="country === 'Isle of Man'" />
  <icon-flagpack-in v-else-if="country === 'India'" />
  <icon-flagpack-io v-else-if="country === 'British Indian Ocean Territory'" />
  <icon-flagpack-iq v-else-if="country === 'Iraq'" />
  <icon-flagpack-ir v-else-if="country === 'Iran'" />
  <icon-flagpack-is v-else-if="country === 'Iceland'" />
  <icon-flagpack-it v-else-if="country === 'Italy'" />
  <icon-flagpack-je v-else-if="country === 'Jersey'" />
  <icon-flagpack-jm v-else-if="country === 'Jamaica'" />
  <icon-flagpack-jo v-else-if="country === 'Jordan'" />
  <icon-flagpack-jp v-else-if="country === 'Japan'" />
  <icon-flagpack-ke v-else-if="country === 'Kenya'" />
  <icon-flagpack-kg v-else-if="country === 'Kyrgyzstan'" />
  <icon-flagpack-kh v-else-if="country === 'Cambodia'" />
  <icon-flagpack-ki v-else-if="country === 'Kiribati'" />
  <icon-flagpack-km v-else-if="country === 'Comoros'" />
  <icon-flagpack-kn v-else-if="country === 'Saint Kitts and Nevis'" />
  <icon-flagpack-kp v-else-if="country === 'North Korea'" />
  <icon-flagpack-kr v-else-if="country === 'South Korea'" />
  <icon-flagpack-kw v-else-if="country === 'Kuwait'" />
  <icon-flagpack-ky v-else-if="country === 'Cayman Islands'" />
  <icon-flagpack-kz v-else-if="country === 'Kazakhstan'" />
  <icon-flagpack-la v-else-if="country === 'Lao People\'s Democratic Republic'" />
  <icon-flagpack-lb v-else-if="country === 'Lebanon'" />
  <icon-flagpack-lc v-else-if="country === 'Saint Lucia'" />
  <icon-flagpack-li v-else-if="country === 'Liechtenstein'" />
  <icon-flagpack-lk v-else-if="country === 'Sri Lanka'" />
  <icon-flagpack-lr v-else-if="country === 'Liberia'" />
  <icon-flagpack-ls v-else-if="country === 'Lesotho'" />
  <icon-flagpack-lt v-else-if="country === 'Lithuania'" />
  <icon-flagpack-lu v-else-if="country === 'Luxembourg'" />
  <icon-flagpack-lv v-else-if="country === 'Latvia'" />
  <icon-flagpack-ly v-else-if="country === 'Libya'" />
  <icon-flagpack-ma v-else-if="country === 'Morocco'" />
  <icon-flagpack-mc v-else-if="country === 'Monaco'" />
  <icon-flagpack-md v-else-if="country === 'Moldova'" />
  <icon-flagpack-me v-else-if="country === 'Montenegro'" />
  <icon-flagpack-mf v-else-if="country === 'Saint Martin'" />
  <icon-flagpack-mg v-else-if="country === 'Madagascar'" />
  <icon-flagpack-mh v-else-if="country === 'Marshall Islands'" />
  <icon-flagpack-mk v-else-if="country === 'Republic of North Macedonia'" />
  <icon-flagpack-ml v-else-if="country === 'Mali'" />
  <icon-flagpack-mm v-else-if="country === 'Myanmar'" />
  <icon-flagpack-mn v-else-if="country === 'Mongolia'" />
  <icon-flagpack-mo v-else-if="country === 'Macao'" />
  <icon-flagpack-mp v-else-if="country === 'Northern Mariana Islands'" />
  <icon-flagpack-mq v-else-if="country === 'Martinique'" />
  <icon-flagpack-mr v-else-if="country === 'Mauritania'" />
  <icon-flagpack-ms v-else-if="country === 'Montserrat'" />
  <icon-flagpack-mt v-else-if="country === 'Malta'" />
  <icon-flagpack-mu v-else-if="country === 'Mauritius'" />
  <icon-flagpack-mv v-else-if="country === 'Maldives'" />
  <icon-flagpack-mw v-else-if="country === 'Malawi'" />
  <icon-flagpack-mx v-else-if="country === 'Mexico'" />
  <icon-flagpack-my v-else-if="country === 'Malaysia'" />
  <icon-flagpack-mz v-else-if="country === 'Mozambique'" />
  <icon-flagpack-na v-else-if="country === 'Namibia'" />
  <icon-flagpack-nc v-else-if="country === 'New Caledonia'" />
  <icon-flagpack-ne v-else-if="country === 'Niger'" />
  <icon-flagpack-nf v-else-if="country === 'Norfolk Island'" />
  <icon-flagpack-ng v-else-if="country === 'Nigeria'" />
  <icon-flagpack-ni v-else-if="country === 'Nicaragua'" />
  <icon-flagpack-nl v-else-if="country === 'Netherlands'" />
  <icon-flagpack-no v-else-if="country === 'Norway'" />
  <icon-flagpack-np v-else-if="country === 'Nepal'" />
  <icon-flagpack-nr v-else-if="country === 'Nauru'" />
  <icon-flagpack-nu v-else-if="country === 'Niue'" />
  <icon-flagpack-nz v-else-if="country === 'New Zealand'" />
  <icon-flagpack-om v-else-if="country === 'Oman'" />
  <icon-flagpack-pa v-else-if="country === 'Panama'" />
  <icon-flagpack-pe v-else-if="country === 'Peru'" />
  <icon-flagpack-pf v-else-if="country === 'French Polynesia'" />
  <icon-flagpack-pg v-else-if="country === 'Papua New Guinea'" />
  <icon-flagpack-ph v-else-if="country === 'Philippines'" />
  <icon-flagpack-pk v-else-if="country === 'Pakistan'" />
  <icon-flagpack-pl v-else-if="country === 'Poland'" />
  <icon-flagpack-pm v-else-if="country === 'Saint Pierre and Miquelon'" />
  <icon-flagpack-pn v-else-if="country === 'Pitcairn'" />
  <icon-flagpack-pr v-else-if="country === 'Puerto Rico'" />
  <icon-flagpack-ps v-else-if="country === 'Palestine'" />
  <icon-flagpack-pt v-else-if="country === 'Portugal'" />
  <icon-flagpack-pw v-else-if="country === 'Palau'" />
  <icon-flagpack-py v-else-if="country === 'Paraguay'" />
  <icon-flagpack-qa v-else-if="country === 'Qatar'" />
  <icon-flagpack-re v-else-if="country === 'Réunion'" />
  <icon-flagpack-ro v-else-if="country === 'Romania'" />
  <icon-flagpack-rs v-else-if="country === 'Serbia'" />
  <icon-flagpack-ru v-else-if="country === 'Russia'" />
  <icon-flagpack-rw v-else-if="country === 'Rwanda'" />
  <icon-flagpack-sa v-else-if="country === 'Saudi Arabia'" />
  <icon-flagpack-sb v-else-if="country === 'Solomon Islands'" />
  <icon-flagpack-sc v-else-if="country === 'Seychelles'" />
  <icon-flagpack-sd v-else-if="country === 'Sudan'" />
  <icon-flagpack-se v-else-if="country === 'Sweden'" />
  <icon-flagpack-sg v-else-if="country === 'Singapore'" />
  <icon-flagpack-sh v-else-if="country === 'Saint Helena, Ascension and Tristan da Cunha'" />
  <icon-flagpack-si v-else-if="country === 'Slovenia'" />
  <icon-flagpack-sj v-else-if="country === 'Svalbard and Jan Mayen'" />
  <icon-flagpack-sk v-else-if="country === 'Slovakia'" />
  <icon-flagpack-sl v-else-if="country === 'Sierra Leone'" />
  <icon-flagpack-sm v-else-if="country === 'San Marino'" />
  <icon-flagpack-sn v-else-if="country === 'Senegal'" />
  <icon-flagpack-so v-else-if="country === 'Somalia'" />
  <icon-flagpack-sr v-else-if="country === 'Suriname'" />
  <icon-flagpack-ss v-else-if="country === 'South Sudan'" />
  <icon-flagpack-st v-else-if="country === 'Sao Tome and Principe'" />
  <icon-flagpack-sv v-else-if="country === 'El Salvador'" />
  <icon-flagpack-sx v-else-if="country === 'Sint Maarten'" />
  <icon-flagpack-sy v-else-if="country === 'Syrian Arab Republic'" />
  <icon-flagpack-sz v-else-if="country === 'Eswatini'" />
  <icon-flagpack-tc v-else-if="country === 'Turks and Caicos Islands'" />
  <icon-flagpack-td v-else-if="country === 'Chad'" />
  <icon-flagpack-tf v-else-if="country === 'French Southern Territories'" />
  <icon-flagpack-tg v-else-if="country === 'Togo'" />
  <icon-flagpack-th v-else-if="country === 'Thailand'" />
  <icon-flagpack-tj v-else-if="country === 'Tajikistan'" />
  <icon-flagpack-tk v-else-if="country === 'Tokelau'" />
  <icon-flagpack-tl v-else-if="country === 'Timor-Leste'" />
  <icon-flagpack-tm v-else-if="country === 'Turkmenistan'" />
  <icon-flagpack-tn v-else-if="country === 'Tunisia'" />
  <icon-flagpack-to v-else-if="country === 'Tonga'" />
  <icon-flagpack-tr v-else-if="country === 'Turkey'" />
  <icon-flagpack-tt v-else-if="country === 'Trinidad and Tobago'" />
  <icon-flagpack-tv v-else-if="country === 'Tuvalu'" />
  <icon-flagpack-tw v-else-if="country === 'Taiwan'" />
  <icon-flagpack-tz v-else-if="country === 'Tanzania'" />
  <icon-flagpack-ua v-else-if="country === 'Ukraine'" />
  <icon-flagpack-ug v-else-if="country === 'Uganda'" />
  <icon-flagpack-um v-else-if="country === 'United States Minor Outlying Islands'" />
  <icon-flagpack-us v-else-if="country === 'United States of America'" />
  <icon-flagpack-uy v-else-if="country === 'Uruguay'" />
  <icon-flagpack-uz v-else-if="country === 'Uzbekistan'" />
  <icon-flagpack-va v-else-if="country === 'Holy See'" />
  <icon-flagpack-vc v-else-if="country === 'Saint Vincent and the Grenadines'" />
  <icon-flagpack-ve v-else-if="country === 'Venezuela'" />
  <icon-flagpack-vg v-else-if="country === 'British Virgin Islands'" />
  <icon-flagpack-vi v-else-if="country === 'Virgin Islands'" />
  <icon-flagpack-vn v-else-if="country === 'Vietnam'" />
  <icon-flagpack-vu v-else-if="country === 'Vanuatu'" />
  <icon-flagpack-wf v-else-if="country === 'Wallis and Futuna'" />
  <icon-flagpack-ws v-else-if="country === 'Samoa'" />
  <icon-flagpack-ye v-else-if="country === 'Yemen'" />
  <icon-flagpack-yt v-else-if="country === 'Mayotte'" />
  <icon-flagpack-za v-else-if="country === 'South Africa'" />
  <icon-flagpack-zm v-else-if="country === 'Zambia'" />
  <icon-flagpack-zw v-else-if="country === 'Zimbabwe'" />
</template>
/* cSpell:enable */
