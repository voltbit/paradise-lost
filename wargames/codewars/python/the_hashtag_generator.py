# https://www.codewars.com/kata/52449b062fb80683ec000024

def the_hashtag_generator(s: str):
    if s == '':
        return False
    hash = '#'
    hash += "".join([w[0].upper() + w[1:].lower() for w in s.split(' ') if w != ''])
    if len(hash) > 140 or len(hash) == 1:
        return False
    return hash


def main():
    assert the_hashtag_generator('') == False, 'Expected an empty string to return False'
    assert the_hashtag_generator('Do We have A Hashtag')[0] == '#', 'Expeted a Hashtag (#) at the beginning.'
    assert the_hashtag_generator('Codewars') == '#Codewars', 'Should handle a single word.'
    assert the_hashtag_generator('Codewars      ') == '#Codewars', 'Should handle trailing whitespace.'
    assert the_hashtag_generator('Codewars Is Nice') == '#CodewarsIsNice', 'Should remove spaces.'
    assert the_hashtag_generator('codewars is nice') == '#CodewarsIsNice', 'Should capitalize first letters of words.'
    assert the_hashtag_generator('CodeWars is nice') == '#CodewarsIsNice', 'Should capitalize all letters of words - all lower case but the first.'
    assert the_hashtag_generator('c i n') == '#CIN', 'Should capitalize first letters of words even when single letters.'
    assert the_hashtag_generator('codewars  is  nice') == '#CodewarsIsNice', 'Should deal with unnecessary middle spaces.'
    assert the_hashtag_generator('Looooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooooong Cat') == False, 'Should return False if the final word is longer than 140 chars.'

if __name__ == '__main__':
    main()

