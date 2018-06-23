curl -X PUT \
  http://localhost:9200/usda \
  -H 'Cache-Control: no-cache' \
  -H 'content-type: application/json' \
  -d '{
  "mappings": {
    "entries": {
    	"properties":{
	    "comm_name": {
          "type": "text"
        },
        "city_name": {
          "type": "text"
        },
        "package_desc": {
          "type": "text"
        },
        "variety_name": {
          "type": "text"
        },
        "subvar_name": {
          "type": "text"
        },
        "location": {
          "type": "geo_point"
        },
        "grade_desc": {
          "type": "text"
        },
        "report_date": {
          "type": "date"
        },
        "low_price_min": {
          "type": "long"
        },
        "high_price_max": {
          "type": "long"
        },
        "mostly_low_min": {
          "type": "long"
        },
        "mostly_high_max": {
          "type": "long"
        },
        "origin_name": {
          "type": "text"
        },
        "district_name": {
          "type": "text"
        },
        "item_size_desc": {
          "type": "text"
        },
        "color": {
          "type": "text"
        },
        "env_desc": {
          "type": "text"
        },
        "unit_sale": {
          "type": "text"
        },
        "quality": {
          "type": "text"
        },
        "condition": {
          "type": "text"
        },
        "appearance": {
          "type": "text"
        },
        "crop": {
          "type": "text"
        },
        "repack": {
          "type": "text"
        },
        "transmode": {
          "type": "text"
        },
        "offerings": {
          "type": "text"
        },
        "market_tone": {
          "type": "text"
        },
        "price_comment": {
          "type": "text"
        },
        "comment": {
          "type": "text"
        },
        "id": {
          "type": "text"
        }
      }
    }
  }
}'
