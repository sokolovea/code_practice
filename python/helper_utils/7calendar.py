import requests

years = range(2026, 2027 + 1)
months = range(1, 12 + 1)
for year in years:
    for month in months:
        calendar_url = f'''https://7calendar.com/view/monthly/pdf/?orientation=l&language=ru&\
                            size=a4&style=0&year={year}&month={month}&weekstart=1&\
                            timeformat=24&download=1&holiday=1&country=RU&weekend=1'''
        response = requests.get(url=calendar_url)
        with open(file=f"{year}_{month:02}.pdf", mode="wb") as f:
            f.write(response.content)
            print(f"Ready for year = {year} and month = {month}")
